package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/hrrydgls/snug/db"
	"github.com/hrrydgls/snug/exceptions"
	"github.com/hrrydgls/snug/models"
	"github.com/hrrydgls/snug/models/responses"
	"gorm.io/gorm"
)

func Email(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	email := r.FormValue("email")

	if email == "" {
		panic(exceptions.Unprocessable("Email is required!"))
	}

	db := db.Connect()

	var user models.User

	result := db.Where("email = ?", email).First(&user)

	if result.Error == nil {
		responses.JSON(w, http.StatusOK, responses.AuthResponse{
			Message: "User is already exist.",
			Result:  "exist",
		})

	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {

		newUser := models.User{
			Email: email,
		}

		if err := db.Create(&newUser).Error; err != nil {
			panic("Could not create user")
		}

		responses.JSON(w, http.StatusCreated, responses.AuthResponse{
			Message: "User created!",
			Result:  "created",
		})

	} else {
		log.Fatal("Database error!", result.Error)
		panic("Database error!")
	}

}
