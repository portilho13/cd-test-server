package main

import (
	"fmt"
	"log"
	"net/http"
)

const API_IP string = "127.0.0.1:8081"

var servers []string

func API(w http.ResponseWriter, r *http.Request) {
	idx := RoundRobin(len(servers))

	url := servers[idx]
	proxyHandler(w, r, url)
}

func ApiRoute(mux *http.ServeMux) {
	mux.HandleFunc("/", API)
}

func InitializeRoutes() http.Handler {
	mux := http.NewServeMux()

	ApiRoute(mux)

	return mux
}

func main() {

	servers = append(servers, "http://127.0.0.1:1337")
	servers = append(servers, "http://127.0.0.1:1338")

	for range 5 {
	}

	mux := InitializeRoutes()
	fmt.Println("Start listening on:", API_IP)
	if err := http.ListenAndServe(API_IP, mux); err != nil {
		log.Fatal(err)
	}

}