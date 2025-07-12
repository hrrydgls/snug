package validators

import (
	"net/mail"

	"github.com/hrrydgls/snug/exceptions"
	"github.com/hrrydgls/snug/requests"
)

func RegisterValidator (registerRequest *requests.RegisterRequest) error {
	if _, error := mail.ParseAddress(registerRequest.Email); error != nil {
		panic(exceptions.Unprocessable("Email is required!"))
	}

	if len(registerRequest.Name) < 3 {
		panic(exceptions.Unprocessable("Name should be at least 3 characters!"))
	}

	if len(registerRequest.Password) < 6 {
		panic(exceptions.Unprocessable("Password should be at least 6 characters!"))
	}

	return nil
}