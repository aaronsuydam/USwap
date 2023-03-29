package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"time"

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

	//END DATABASE CODE
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
	query := `CREATE TABLE IF NOT EXISTS items(item_id text, item_name text, item_description text, user_id int, image_path text)`
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
	query := `CREATE TABLE IF NOT EXISTS swap(swap_id text, sender_id int, sender_item_id int, receiver_id int, receiver_item_id int)`
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
func CreateUser(userName string, userEmail string, userPassword string) error {

	// Generate a user ID
	userID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO users (user_id, user_name, user_email, user_password) VALUES (?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, userID, userName, userEmail, userPassword)
	return err
}

// Add an item to the items table upon user listing the item
func CreateItem(itemName string, itemDescription string, userID string, imagePath string) error {

	// Generate an item ID
	itemID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO items (item_id int, item_name text, item_description text, user_id int, image_path text) VALUES (?, ?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, itemID, itemName, itemDescription, userID, imagePath)
	return err
}

// Add a swap request into swap table
func CreateSwapRequest(senderID string, senderItemID string, receiverID string, receiverItemID string) error {
	// Generate an item ID
	swapID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO swap (swap_id text, sender_id text, sender_item_id text, receiver_id text, receiver_item_id) VALUES (?, ?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, swapID, senderID, senderItemID, receiverID, receiverItemID)
	return err
}

// Complete these
func GetUser(userID string) {
	temp := 1
	temp += 1
}
func GetItem(itemID string) {
	temp := 1
	temp += 1
}

func GetSwapRequest(swapID string) {
	temp := 1
	temp += 1
}
