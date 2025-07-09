package routes

import (
	"net/http"

	"github.com/hrrydgls/snug/handlers"
	"github.com/hrrydgls/snug/handlers/auth"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/auth/email", auth.Email)
	mux.HandleFunc("/login", handlers.LoginHandler)

	return mux
}
