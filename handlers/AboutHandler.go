package handlers

import (
	"net/http"

	"github.com/hrrydgls/snug/models/responses"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, responses.Response{
		Message: "Welcome to the About Page!",
	})
}
