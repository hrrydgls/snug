package auth

import (
	"encoding/json"
	"net/http"

	"github.com/hrrydgls/snug/db"
	"github.com/hrrydgls/snug/models"
	"gorm.io/gorm"

	"github.com/hrrydgls/snug/models/responses"
)

func Email(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseForm()
	email := r.FormValue("email")

	if email == "" {
		validationError := responses.Response{
			Code:    422,
			Message: "Invalid email address",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(validationError)
		return
	}

	db := db.Connect()

	var user models.User

	result := db.Where("email = ?", email).First(&user)

	if result.Error == nil {
		// User exists
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User already exists",
		})
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		// DB error
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Database error",
		})
		return
	}

	// User does not exist → Create user
	newUser := models.User{
		Email: email,
	}

	if err := db.Create(&newUser).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Could not create user",
		})
		return
	}

	// Return created user as JSON
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
