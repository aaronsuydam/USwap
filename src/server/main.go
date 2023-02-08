package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/atxfjrotc/uswap/src/server/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "HELLO WORLD",
	}

	log.Fatal(srv.ListenAndServe())
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		LoginSuccess string `json:"loginSuccess"`
	}{
		LoginSuccess: "False",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func main() {

	// Routes
	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorldHandler)
	r.HandleFunc("/hello-world", HelloWorldHandler)
	r.HandleFunc("/login", LoginPost).Methods("POST")
	r.HandleFunc("/login", LoginPost)

	srv := &http.Server{
		Handler:      handlers.CORS()(r),
		Addr:         "127.0.0.1:4201",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
