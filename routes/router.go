package routes

import (
	"net/http"

	"github.com/hrrydgls/snug/handlers"
	"github.com/hrrydgls/snug/handlers/auth"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {

	authHandler := &auth.Handler{DB: db}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/auth/email", authHandler.Email)
	mux.HandleFunc("/auth/register", authHandler.Register)
	mux.HandleFunc("/login", handlers.LoginHandler)

	return mux
}
