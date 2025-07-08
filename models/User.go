package models

type User struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	EmailVerifiedAt bool `json:"email_verified_at"`
	Password string `json:"password"`
}