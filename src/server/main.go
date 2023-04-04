package main

import (
	"log"
	"net/http"
	"time"

	"github.com/atxfjrotc/uswap/src/server/handler"
	_ "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "fmt"
	_ "text/template"
)

func main() {

	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)

	r.HandleFunc("/login", handler.LoginPost).Methods("POST")
	r.HandleFunc("/signup", handler.SignUpPost).Methods("POST")

	r.HandleFunc("/item", handler.GetItem).Methods("GET")
	r.HandleFunc("/item/create", handler.CreateListing).Methods("POST")
	//r.HandleFunc("/item/delete", handler.DeleteListing).Methods("POST") // Delete an item
	//r.HandleFunc("/item/modify", handler.ModifyItem).Methods("PUT") // Modify an item

	r.HandleFunc("/swap", handler.AcceptSwapRequest).Methods("GET")
	r.HandleFunc("/swap/create", handler.AcceptSwapRequest).Methods("GET")
	r.HandleFunc("/swap/accept", handler.AcceptSwapRequest).Methods("POST")
	r.HandleFunc("/swap/reject", handler.AcceptSwapRequest).Methods("POST")

	srv := &http.Server{
		Addr:         ":4201",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
