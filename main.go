package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hrrydgls/snug/routes"
	"github.com/hrrydgls/snug/middlewares"
)

func main() {
	loadEnv()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8888"
	}

	router := routes.SetupRoutes()
	handler := middlewares.RecoveryMiddleware(router)

	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env variables")
	}
}
