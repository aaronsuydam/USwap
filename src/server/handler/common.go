package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	// "github.com/joho/godotenv"
	"github.com/atxfjrotc/uswap/src/server/db"
	"github.com/atxfjrotc/uswap/src/server/utils"
	"github.com/golang-jwt/jwt/v5"
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

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey = []byte(os.Getenv("RSA_PRIVATE_KEY"))

func LoginPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var login Login
	json.Unmarshal(body, &login)

	ctx := db.Ctx
	er := db.DB.PingContext(ctx)
	if er != nil {
		panic(er)
	}

	tsql := fmt.Sprintf("SELECT Password FROM TestSchema.Users WHERE Name = @Name")
	
	rows, err := db.DB.QueryContext(ctx, tsql, sql.Named("Name", login.Username))
	if err != nil {
		fmt.Println("Error with creating db query")
		panic(err)
	}
	defer rows.Close()

	// var hash string
	var pwd string
	for rows.Next() {
		err := rows.Scan(&pwd)
		if err != nil {
			log.Fatal(err)
		}
	}

	// success := utils.CheckPasswordHash(login.Password, string(pwd))

	success := true
	if !success {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute)

	claims := &Claims{
		Username: login.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	loginJson, err := json.Marshal(login)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(loginJson)
}

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpPost(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var signup SignUp
	json.Unmarshal(body, &signup)

	createID, err := db.CreateUser(signup.Username, signup.Email, signup.Password)
	if err != nil {
		log.Fatal("Error creating User: ", err.Error())
	}
	fmt.Printf("Inserted ID: %d successfully.\n", createID)

	w.WriteHeader(http.StatusOK)
}

type Item struct {
	Name        string `json:"itemName"`
	Description string `json:"itemDescription"`
	UserID      string `json:"userID"`
	ImagePath   string `json:"imagePath"`
}

func CreateListing(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
	enableCors(&w)

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

func SearchItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	type Search struct {
		Search string `json:"search"`
	}
	var search Search
	json.Unmarshal(body, &search)

}

type Swap struct {
	SenderID       string `json:"senderID"`
	SenderItemID   string `json:"senderItemID"`
	ReceiverID     string `json:"receiverID"`
	ReceiverItemID string `json:"receiverItemID"`
}

func CreateSwapRequest(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

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
	enableCors(&w)

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
	enableCors(&w)

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
	enableCors(&w)

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
