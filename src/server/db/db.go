package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/atxfjrotc/uswap/src/server/utils"
	uuid "github.com/nu7hatch/gouuid"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/microsoft/go-mssqldb"
)

// Global variable to hold DB connection
var DB *sql.DB
var Ctx = context.Background()
var username string
var password string
var hostname string
var dbname string
var port int
var portS string

func Initialize() (err error) {
	// Establish Database Connection
	username = os.Getenv("DBUSERNAME")
	password = os.Getenv("DBPASSWORD")
	hostname = os.Getenv("DBHOSTNAME")
	dbname = os.Getenv("DBNAME")
	portS = os.Getenv("PORT")
	port, err = strconv.Atoi(portS)
	DB, err = dbConnection()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// err = createUserTable()
	// if err != nil {
	// 	log.Printf("Create user table failed with error %s", err)
	// 	return err
	// }
	// err = createItemsTable()
	// if err != nil {
	// 	log.Printf("Create userItems table failed with error %s", err)
	// 	return err
	// }
	// err = createSwapTable()
	// if err != nil {
	// 	log.Printf("Create swap table failed with error %s", err)
	// 	return err
	// }
	return err
}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func dbConnection() (*sql.DB, error) {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		hostname, username, password, port, dbname)
	var err error
	// Create connection pool
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = DB.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")

	count, err := ReadUsers()
    if err != nil {
        log.Fatal("Error reading Users: ", err.Error())
    }
    fmt.Printf("Read %d row(s) successfully.\n", count)

	return DB, nil
}

// User table maintains all users
func createUserTable() error {
	var err error

	err = DB.PingContext(Ctx); if err != nil {
		log.Printf("Error %s when creating users table", err)
		return err
	}

	queryStatement := `CREATE TABLE IF NOT EXISTS users(user_id text, user_name text, user_email text, user_password text)`

	query, err := DB.Prepare(queryStatement); if err != nil {
		return err
	}
	defer query.Close()

	return nil
}

// Items table maintains all actively listed items
func createItemsTable() error {
	var err error

	err = DB.PingContext(Ctx); if err != nil {
		log.Printf("Error %s when creating items table", err)
		return err
	}

	queryStatement := `CREATE TABLE IF NOT EXISTS items(item_id text, item_name text, item_description text, user_id text, image_path text)`

	query, err := DB.Prepare(queryStatement); if err != nil {
		return err
	}
	defer query.Close()

	return nil
}

// Swap table maintains all active swap requests
func createSwapTable() error {
	var err error

	err = DB.PingContext(Ctx); if err != nil {
		log.Printf("Error %s when creating swap table", err)
		return err
	}

	queryStatement := `CREATE TABLE IF NOT EXISTS swap(swap_id text, sender_id text, sender_item_id text, receiver_id text, receiver_item_id text)`

	query, err := DB.Prepare(queryStatement); if err != nil {
		return err
	}
	defer query.Close()

	return nil
}

// Add a user to the user table on signup
func CreateUser(userName string, userEmail string, userPassword string) (int64, error) {
	var err error

	// Generate a user ID
	// byteUserID, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// userID = byteUserID.String()

	if DB == nil {
		err = errors.New("CreateUser: db is null")
		return -1, err
	}

	queryStatement := `
		INSERT INTO TestSchema.Users (Name, Email, Password) VALUES (@Name, @Email, @Password);
		select isNull(SCOPE_IDENTITY(), -1);
	`

	query, err := DB.Prepare(queryStatement); if err != nil {
		return -1, err
	}
	defer query.Close()

	userPassword, _ = utils.HashPassword(userPassword)

	row := query.QueryRowContext(
		Ctx,
		sql.Named("Name", userName),
		sql.Named("Email", userEmail),
		sql.Named("Password", userPassword))
	
	var newID int64
	err = row.Scan(&newID); if err != nil {
		return -1, err
	}

	return newID, nil
}

func ReadUsers() (int, error) {
	err := DB.PingContext(Ctx)
    if err != nil {
        return -1, err
    }

    tsql := fmt.Sprintf("SELECT Id, Name, Email, Password FROM TestSchema.Users;")

    // Execute query
    rows, err := DB.QueryContext(Ctx, tsql)
    if err != nil {
        return -1, err
    }

    defer rows.Close()

    var count int

    // Iterate through the result set.
    for rows.Next() {
        var name, email, password string
        var id int

        // Get values from row.
        err := rows.Scan(&id, &name, &email, &password)
        if err != nil {
            return -1, err
        }

        fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s\n", id, name, email, password)
        count++
    }

    return count, nil
}

// Add an item to the items table upon user listing the item
func CreateItem(itemName string, itemDescription string, userID string, imagePath string) (itemID string, err error) {

	// Generate an item ID
	byteItemID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	itemID = byteItemID.String()

	query := `INSERT INTO items (item_id, item_name, item_description, user_id, image_path) VALUES (?,?,?,?,?)`
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

	query := `INSERT INTO swap (swap_id, sender_id, sender_item_id, receiver_id, receiver_item_id) VALUES (?,?,?,?,?)`
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

/*
func searchItems(search string) ([]Item, err) {
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
}*/

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
