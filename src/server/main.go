package main

import (
	"log"
	"net/http"
	"time"

	"github.com/atxfjrotc/uswap/src/server/db"
	"github.com/atxfjrotc/uswap/src/server/handler"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"

	_ "fmt"
	_ "text/template"
)

func main() {
	// Environment Vars
	err := godotenv.Load("src/server/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = db.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)

	r.HandleFunc("/login", handler.LoginPost).Methods("POST", "OPTIONS")
	r.HandleFunc("/signup", handler.SignUpPost).Methods("POST", "OPTIONS")

	r.HandleFunc("/item", handler.GetItem).Methods("GET")
	r.HandleFunc("/item/create", handler.CreateListing).Methods("POST", "OPTIONS")
	//r.HandleFunc("/item/delete", handler.DeleteListing).Methods("POST") // Delete an item
	//r.HandleFunc("/item/modify", handler.ModifyItem).Methods("PUT") // Modify an item

	r.HandleFunc("/swap", handler.AcceptSwapRequest).Methods("GET")
	r.HandleFunc("/swap/create", handler.AcceptSwapRequest).Methods("GET", "OPTIONS")
	r.HandleFunc("/swap/accept", handler.AcceptSwapRequest).Methods("POST", "OPTIONS")
	r.HandleFunc("/swap/reject", handler.AcceptSwapRequest).Methods("POST", "OPTIONS")

	srv := &http.Server{
		Addr:         ":4201",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
