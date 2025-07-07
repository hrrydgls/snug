package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/hrrydgls/snug/models/responses"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	response := responses.Response{
		Code: 200,
		Message: "Everything looks allright!",
	}
	json.NewEncoder(w).Encode(response)
}	