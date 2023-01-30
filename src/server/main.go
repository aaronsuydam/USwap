package main

import (
	"fmt"
	"net/http"

	"github.com/atxfjrotc/uswap/src/server/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Page Accessed")
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "HELLO WORLD",
	}

	jsonBytes, err := utils.StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Testlog")
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "Login",
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
	r.HandleFunc("/login", LoginHandler)

	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:4201",
	// 	//Addr:         ":" + os.Getenv("PORT"),
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	http.ListenAndServe("127.0.0.1:4201", handlers.CORS()(r))

	// log.Fatal(srv.ListenAndServe())
}
