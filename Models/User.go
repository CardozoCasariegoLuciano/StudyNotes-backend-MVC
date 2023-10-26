package models

type User struct {
	Id       int    `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    string `json:"image"`
	Role     string `json:"role"`
}
