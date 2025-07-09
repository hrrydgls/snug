package handlers

import (
	"net/http"
	"github.com/hrrydgls/snug/models/responses"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, responses.Response{
		Message: "Everything looks allright!",
	})
}	