package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/atxfjrotc/uswap/src/server/db"
	"github.com/atxfjrotc/uswap/src/server/utils"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
			w.Header().Set("Content-Type", "application/json")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
	enableCors(&w)
	w.Write(jsonBytes)
}

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var signup SignUp
	json.Unmarshal(body, &signup)

	signup.Password, _ = utils.HashPassword(string(signup.Password)) // Hash password
	err = db.CreateUser(signup.Username, signup.Email, signup.Password)
	if err != nil {
		log.Fatal("Failed to sign up user")
	}

	enableCors(&w)
}

type Item struct {
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
	UserID      string `json:"userID"`
	ImagePath   string `json:"imagePath"`
}

func CreateListing(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var item Item
	json.Unmarshal(body, &item)

	err = db.CreateItem(item.Name, item.Description, item.UserID, item.ImagePath)
	if err != nil {
		log.Fatal("Failed to create item listing")
	}
}

type Swap struct {
	SenderID       string `json:"senderID"`
	SenderItemID   string `json:"senderItemID"`
	ReceiverID     string `json:"receiverID"`
	ReceiverItemID string `json:"receiverItemID"`
}

func CreateSwapRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var swap Swap
	json.Unmarshal(body, &swap)

	err = db.CreateItem(swap.SenderID, swap.SenderItemID, swap.ReceiverID, swap.ReceiverItemID)
	if err != nil {
		log.Fatal("Failed to create the swap request")
	}
}
