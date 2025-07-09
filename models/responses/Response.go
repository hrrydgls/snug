package responses

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

type Response struct {
	Message string `json:"message"`
}

type AuthResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
}

