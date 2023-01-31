package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/atxfjrotc/uswap/src/server/utils"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "uswap"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {
	// r := mux.NewRouter()

	// r.HandleFunc("/hello-world", helloWorld)

	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)

	db.Close()
	db, err = sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	//CREATE TABLE IF NOT EXISTS users(student_id int primary key, user_name text, user_email text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)

	// log.Fatal(srv.ListenAndServe())
	fmt.Println("Hello, World!")
	utils.Fun()
}

// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	var data = struct {
// 		Title string `json:"title"`
// 	}{
// 		Title: "Golang + Angular Starter Kit",
// 	}

// 	jsonBytes, err := utils.StructToJSON(data)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonBytes)
// 	return
// }
