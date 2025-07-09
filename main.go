package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hrrydgls/snug/handlers"
	"github.com/hrrydgls/snug/handlers/auth"
	"github.com/hrrydgls/snug/middlewares"
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

	mux.HandleFunc("/auth/email", auth.Email)

	mux.HandleFunc("/login", handlers.LoginHandler)

	log.Fatal(http.ListenAndServe(":"+port, middlewares.RecoveryMiddleware(mux)))

}