package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/atxfjrotc/uswap/src/server/db"
	"github.com/atxfjrotc/uswap/src/server/utils"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type user struct {
	userId       int
	userName     string
	userEmail    string
	userPassword string
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var login Login
	json.Unmarshal(body, &login)

	rows, err := db.DB.Query("SELECT user_password FROM users2 WHERE user_name = ?", login.Username)
	if err != nil {
		fmt.Println("Error with creating db query")
		log.Fatal(err)
	}

	defer rows.Close()
	var hash string
	for rows.Next() {
		if err := rows.Scan(&hash); err != nil {
			log.Fatal(err)
		}
	}

	success := utils.CheckPasswordHash(login.Password, string(hash))

	var data struct {
		LoginSuccess string `json:"loginSuccess"`
	}
	if success {
		data.LoginSuccess = `True`
	} else {
		data.LoginSuccess = `False`
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var login Login
	json.Unmarshal(body, &login)

	hash, _ := utils.HashPassword(string(login.Password))
	rows, err := db.DB.Query("SELECT COUNT(*) as count FROM users2")
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
	stmt, err := db.DB.PrepareContext(ctx, query)
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
}
