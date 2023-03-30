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
	_, err = db.CreateUser(signup.Username, signup.Email, signup.Password)
	if err != nil {
		log.Fatal("Failed to sign up user")
	}
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

	_, err = db.CreateItem(item.Name, item.Description, item.UserID, item.ImagePath)
	if err != nil {
		log.Fatal("Failed to create item listing")
	}
}

type ItemID struct {
	ItemID string `json:"itemID"`
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var itemID ItemID
	json.Unmarshal(body, &itemID)

	item, err := db.GetItem(itemID.ItemID)

	jsonBytes, err := utils.StructToJSON(item)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
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

	_, err = db.CreateItem(swap.SenderID, swap.SenderItemID, swap.ReceiverID, swap.ReceiverItemID)
	if err != nil {
		log.Fatal("Failed to create the swap request")
	}
}

type SwapID struct {
	SwapID string `json:"swapID"`
}

func GetSwapRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var swapID SwapID
	json.Unmarshal(body, &swapID)

	swapRequest, err := db.GetItem(swapID.SwapID)

	jsonBytes, err := utils.StructToJSON(swapRequest)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func AcceptSwapRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var swapID SwapID
	json.Unmarshal(body, &swapID)

	err = db.AcceptSwapRequest(swapID.SwapID)
	if err != nil {
		log.Panic("Failed to accept swap request")
	}
}

func RejectSwapRequest(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var swapID SwapID
	json.Unmarshal(body, &swapID)

	err = db.RejectSwapRequest(swapID.SwapID)
	if err != nil {
		log.Panic("Failed to accept swap request")
	}
}
