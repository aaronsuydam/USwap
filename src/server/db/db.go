package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	uuid "github.com/nu7hatch/gouuid"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Move to environment variables
	username = "sql9600821"
	password = "pbK4HDuGPn"
	hostname = "sql9.freesqldatabase.com"
	dbname   = "sql9600821"
)

// Global variable to hold DB connection
var DB *sql.DB

func init() {
	// Establish Database Connection
	var err error
	DB, err = dbConnection()
	if err != nil {
		return
	}
	log.Printf("Successfully connected to database")
	err = createUserTable()
	if err != nil {
		log.Printf("Create user table failed with error %s", err)
		return
	}
	err = createItemsTable()
	if err != nil {
		log.Printf("Create userItems table failed with error %s", err)
		return
	}
	err = createSwapTable()
	if err != nil {
		log.Printf("Create swap table failed with error %s", err)
		return
	}
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func dbConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return nil, err
	}
	//defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	//defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return nil, err
	}
	return db, nil
}

// User table maintains all users
func createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users(user_id text, user_name text,user_email text,
        user_password text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}
	return nil
}

// Items table maintains all actively listed items
func createItemsTable() error {
	query := `CREATE TABLE IF NOT EXISTS items(item_id text, item_name text, item_description text, user_id text, image_path text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}
	return nil
}

// Swap table maintains all active swap requests
func createSwapTable() error {
	query := `CREATE TABLE IF NOT EXISTS swap(swap_id text, sender_id text, sender_item_id text, receiver_id text, receiver_item_id text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := DB.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}
	return nil
}

// Add a user to the user table on signup
func CreateUser(userName string, userEmail string, userPassword string) (userID string, err error) {

	// Generate a user ID
	byteUserID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	userID = byteUserID.String()

	query := `INSERT INTO users (user_id, user_name, user_email, user_password) VALUES (?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, userName, userEmail, userPassword)
	if err != nil {
		log.Fatal(err)
	}
	return userID, err
}

// Add an item to the items table upon user listing the item
func CreateItem(itemName string, itemDescription string, userID string, imagePath string) (itemID string, err error) {

	// Generate an item ID
	byteItemID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	itemID = byteItemID.String()

	query := `INSERT INTO items (item_id text, item_name text, item_description text, user_id text, image_path text) VALUES (?, ?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, itemID, itemName, itemDescription, userID, imagePath)
	return itemID, err
}

// Add a swap request into swap table
func CreateSwapRequest(senderID string, senderItemID string, receiverID string, receiverItemID string) (swapID string, err error) {
	// Generate an item ID
	byteSwapID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	swapID = byteSwapID.String()

	query := `INSERT INTO swap (swap_id text, sender_id text, sender_item_id text, receiver_id text, receiver_item_id text) VALUES (?, ?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, swapID, senderID, senderItemID, receiverID, receiverItemID)
	if err != nil {
		log.Fatal(err)
	}
	return swapID, err
}

type User struct {
	user_id       string
	user_name     string
	user_email    string
	user_password string
}

func GetUser(userID string) (User, error) {

	var user User

	row := DB.QueryRow("SELECT * FROM users WHERE user_id = ?", userID)
	if err := row.Scan(&user.user_id, &user.user_name, &user.user_email, &user.user_password); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("userById %v: no such user", userID)
		}
		return user, fmt.Errorf("userById %v: %v", userID, err)
	}

	return user, nil
}

type Item struct {
	item_id          string
	item_name        string
	item_description string
	user_id          string
	image_path       string
}

func GetItem(itemID string) (Item, error) {
	var item Item

	row := DB.QueryRow("SELECT * FROM items WHERE item_id = ?", itemID)
	if err := row.Scan(&item.item_id, &item.item_name, &item.item_description, &item.user_id, &item.image_path); err != nil {
		if err == sql.ErrNoRows {
			return item, fmt.Errorf("getItem %v: no such item", itemID)
		}
		return item, fmt.Errorf("getItem %v: %v", itemID, err)
	}

	return item, nil
}

type Swap struct {
	swap_id          string
	sender_id        string
	sender_item_id   string
	receiver_id      string
	receiver_item_id string
}

func GetSwapRequest(swapID string) (Swap, error) {
	var swap Swap

	row := DB.QueryRow("SELECT * FROM swap WHERE swap_id = ? ", swapID)
	if err := row.Scan(&swap.swap_id, &swap.sender_id, &swap.sender_item_id, &swap.receiver_id, &swap.receiver_item_id); err != nil {
		if err == sql.ErrNoRows {
			return swap, fmt.Errorf("getItem %v: no such swap", swapID)
		}
		return swap, fmt.Errorf("getItem %v: %v", swapID, err)
	}

	return swap, nil
}

func GetUserItems(userID string) ([]Item, error) {
	var items []Item

	rows, err := DB.Query("SELECT * FROM items WHERE user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("alluserItems %v: %v", userID, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.item_id, &item.item_name, &item.item_description, &item.user_id, &item.image_path); err != nil {
			return nil, fmt.Errorf("alluserItems %v: %v", userID, err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("alluserItems %v: %v", userID, err)
	}
	return items, nil

}

func AcceptSwapRequest(swapID string) (err error) {
	swap, err := GetSwapRequest(swapID)
	if err != nil {
		log.Fatalf("Swap ID %v not found in swap table", swapID)
		return err
	}
	senderItem, err := GetItem(swap.sender_item_id)
	if err != nil {
		log.Fatalf("Item ID %v not found in items table", swap.sender_item_id)
	}
	receiverItem, err := GetItem(swap.receiver_item_id)
	if err != nil {
		log.Fatalf("Item ID %v not found in items table", swap.receiver_item_id)
	}

	query := `UPDATE items SET user_id=? WHERE item_id=?`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Swap user IDs associated with the items
	_, err = stmt.ExecContext(ctx, receiverItem.user_id, senderItem.item_id)
	if err != nil {
		log.Fatal("Updating items table failed")
	}
	_, err = stmt.ExecContext(ctx, senderItem.user_id, receiverItem.item_id)
	if err != nil {
		log.Fatal("Updating items table failed")
	}

	return err
}

func RejectSwapRequest(swapID string) (err error) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = DB.ExecContext(ctx, "DELETE FROM swap WHERE swap_id = ?", swapID)
	if err != nil {
		log.Fatalf("Failed to reject Swap Request with id %v", swapID)
	}
	return err
}
