package main

import (
	"log"
	"net/http"
	"time"

	"github.com/atxfjrotc/uswap/src/server/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "fmt"
	_ "text/template"
)

func main() {

	r := mux.NewRouter()
	r.Use(handler.CorsMiddleware)
	r.HandleFunc("/login", handler.LoginPost).Methods("POST")
	r.HandleFunc("/login", handler.LoginPost)
	r.HandleFunc("/signup", handler.SignUpPost).Methods("POST")

	srv := &http.Server{
		Addr:         ":4201",
		Handler:	  r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
