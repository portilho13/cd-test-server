package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const API_IP string = "127.0.0.1:1338"

func TestAPI(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "This is server 2!"})
}

func TestApiRoute(mux *http.ServeMux) {
	mux.HandleFunc("/", TestAPI)
}

func InitializeRoutes() http.Handler {
	mux := http.NewServeMux()

	TestApiRoute(mux)

	return mux
}

func main() {

	mux := InitializeRoutes()
	fmt.Println("Start listening on:", API_IP)
	if err := http.ListenAndServe(API_IP, mux); err != nil {
		log.Fatal(err)
	}

}