package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hrrydgls/snug/db"
	"github.com/hrrydgls/snug/middlewares"
	"github.com/hrrydgls/snug/routes"
)

func main() {
	loadEnv()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8888"
	}

	appEnv := os.Getenv("APP_ENV")

	database := db.Connect()

	router := routes.SetupRoutes(database)
	handler := middlewares.RecoveryMiddleware(router)

	log.Printf("Server is running on port '%s' and the env is '%s'", port, appEnv)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env variables")
	}
}
