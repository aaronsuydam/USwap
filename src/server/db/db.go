package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	err = createUserTable(DB)
	if err != nil {
		log.Printf("Create user table failed with error %s", err)
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

func createUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users2(user_id int, user_name text,user_email text,
        user_password text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}
	return nil
}
func createUserItemsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS userItems(user_items text, user_id int)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}
	return nil
}

func addItems(id int, item string, db *sql.DB) {
	query := `INSERT INTO usersItems (user_items text, user_id int) VALUES (?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, item, id) //exec to create usertable

}

func Test() {
	err := DB.Ping()
	if err != nil {
		fmt.Println("Error pinging from test")
	} else {
		fmt.Println("Successfully pinging from test")
	}
}
