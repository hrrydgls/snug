package models 

type Language struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Active bool `json:"active"`
	CreatedAt string `json:"created_at"`
}