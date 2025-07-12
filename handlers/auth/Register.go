package auth

import (
	"net/http"

	"github.com/hrrydgls/snug/requests"
	"github.com/hrrydgls/snug/validators"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	email := r.FormValue("email")
	name := r.FormValue("name")
	password := r.FormValue("password")


	registerRequest := requests.RegisterRequest {
		Email: email,
		Name: name,
		Password: password,
	}

	validators.RegisterValidator(&registerRequest)

}