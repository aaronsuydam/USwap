package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"time"

	_ "text/template"

	"github.com/atxfjrotc/uswap/src/server/utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "sql9595555"
	password = "xXEPeR9JDl"
	hostname = "sql9.freesqldatabase.com"
	dbname   = "sql9595555"
)

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
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return nil, err
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return nil, err
	}
	log.Printf("rows affected %d\n", no)

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
	log.Printf("Connected to DB %s successfully\n", dbname)
	return db, nil
}

func createUserTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS users2(user_id int, user_name text,user_email text,
        user_password text)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}

	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}

type user struct {
	userId int
	userN  string
	userE  string
	userP  string
}

func manuallyAdd(db *sql.DB) error {

	u1 := user{
		userId: 1234568790,
		userN:  "Test user",
		userE:  "testuser@test.com",
		userP:  "password",
	}

	query := `INSERT INTO users2 (user_id, user_name, user_email, user_password) VALUES (?, ?, ?, ?)`
	//query := "INSERT INTO product(product_name, product_price) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, u1.userId, u1.userN, u1.userE, u1.userP)
	if err != nil {
		log.Printf("Error %s when inserting row into user table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)

	return nil
}
func main() {

	//BEGIN DATABASE CODE
	db, err := dbConnection()
	if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}
	defer db.Close()
	log.Printf("Successfully connected to database")
	err = createUserTable(db)
	if err != nil {
		log.Printf("Create user table failed with error %s", err)
		return
	}

	err = manuallyAdd(db)
	if err != nil {
		log.Printf("Manually add user FAILED with error %s", err)
		return
	}

	//END DATABASE CODE

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Golang + Angular Starter Kit",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
	return
}
