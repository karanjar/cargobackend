package main

import (
	"fmt"
	"net/http"

	"github.com/karanjar/cargobackend.git/config"
	"github.com/karanjar/cargobackend.git/handlers"
	"github.com/karanjar/cargobackend.git/middleware"
)

func main() {
	config.ConnectDb()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Carhandler)
	mux.HandleFunc("/cars/", handlers.Carhandler)

	wrappedMux := middleware.Logger(mux)
	wrappedMux = middleware.SecureHeaders(wrappedMux)

	fmt.Println("HTTP server listenning on port 8080")
	http.ListenAndServe(":8080", wrappedMux)
}
