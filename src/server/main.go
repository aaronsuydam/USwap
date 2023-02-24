package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "text/template"

	"github.com/atxfjrotc/uswap/src/server/utils"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

const (
	// Move to environment variables
	username = "sql9595555"
	password = "xXEPeR9JDl"
	hostname = "sql9.freesqldatabase.com"
	dbname   = "sql9595555"
)

var db *sql.DB

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
	userId       int
	userName     string
	userEmail    string
	userPassword string
}

type Person struct {
	Name string `json:"Name"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "HELLO WORLD",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Hashing functionality
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var login Login
	json.Unmarshal(body, &login)
	fmt.Println(login.Username)

	hash, _ := HashPassword(string(login.Password))
	rows, err := db.Query("SELECT COUNT(*) as count FROM users2")
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for rows.Next() {
		rows.Scan(&count)
	}

	// Creates ID for next user in database
	u1 := user{
		userId:       count + 1,
		userName:     login.Username,
		userEmail:    "testuser@test.com",
		userPassword: hash,
	}

	query := `INSERT INTO users2 (user_id, user_name, user_email, user_password) VALUES (?, ?, ?, ?)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u1.userId, u1.userName, u1.userEmail, u1.userPassword)

	if err != nil {
		log.Printf("Error %s when inserting row into user table", err)
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var login Login
	json.Unmarshal(body, &login)

	// Test user database named user2
	rows, err := db.Query("SELECT user_password FROM users2 WHERE user_name = ?", login.Username)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var hash string
	for rows.Next() {
		if err := rows.Scan(&hash); err != nil {
			log.Fatal(err)
		}
	}

	success := CheckPasswordHash(login.Password, string(hash))

	var data struct {
		LoginSuccess string `json:"loginSuccess"`
	}
	if success {
		data.LoginSuccess = `True`
	} else {
		data.LoginSuccess = `False`
	}

	// Just for testing
	fmt.Println(data.LoginSuccess)

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func TestEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	w.Write([]byte("foo"))

	// p1 := &Person{}
	// utils.GetJson("http://localhost:4201/test", p1)
	// println(p1.Email)

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	fmt.Println(w.Header())
}

func main() {

	// Establish Database Connection
	db, _ = dbConnection()
	/*if err != nil {
		log.Printf("Error %s when getting db connection", err)
		return
	}*/
	defer db.Close()
	log.Printf("Successfully connected to database")
	err := createUserTable(db)
	if err != nil {
		log.Printf("Create user table failed with error %s", err)
		return
	}
	//CREATE TABLE IF NOT EXISTS users(student_id int primary key, user_name text, user_email text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)
	//END DATABASE CODE

	// Routes
	r := mux.NewRouter()

	// r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/", HelloWorldHandler).Methods("GET", "POST")
	r.HandleFunc("/login", LoginPost).Methods("POST")
	r.HandleFunc("/login", LoginPost)
	r.HandleFunc("/signup", SignUpPost).Methods("GET", "POST")
	r.HandleFunc("/test", TestEndpoint).Methods(http.MethodGet, http.MethodPost)

	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Handler:	  handlers.CORS()(r),
		Addr:         "127.0.0.1:4201",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
