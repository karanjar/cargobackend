package main

import (
	"fmt"
	"net/http"

	"github.com/karanjar/cargobackend.git/config"
	"github.com/karanjar/cargobackend.git/handlers"
)

func main() {
	config.ConnectDb()

	http.HandleFunc("/", handlers.Carhandler)
	http.HandleFunc("/cars/", handlers.Carhandler)

	fmt.Println("HTTP server listenning on port 8080")
	http.ListenAndServe(":8081", nil)
}
