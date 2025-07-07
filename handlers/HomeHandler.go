package handlers

import (
	"net/http"
	"github.com/hrrydgls/snug/models"
	"encoding/json"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	setting := models.Setting{
		Title:       "Welcome to Snug",
		Description: "A simple and cozy web application",
		Keywords:    "snug, cozy, web application",
	}

	json.NewEncoder(w).Encode(setting)
}