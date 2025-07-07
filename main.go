package main

import (
	"net/http"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/hrrydgls/snug/handlers"
)

func main () {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8888"
	}


	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)

	mux.HandleFunc("/about", handlers.AboutHandler)

	mux.HandleFunc("/login", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(":"+port, mux))

}