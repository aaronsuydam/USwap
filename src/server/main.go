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
	r.HandleFunc("/login", handler.LoginPost).Methods("POST")
	r.HandleFunc("/login", handler.LoginPost)
	r.HandleFunc("/signup", handler.SignUpPost).Methods("POST")

	srv := &http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         "127.0.0.1:4201",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
