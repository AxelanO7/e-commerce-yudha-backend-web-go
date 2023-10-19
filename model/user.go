package model

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Name     string `json:"name_user"`
	Address  string `json:"address_user"`
	Phone    string `json:"phone_user"`
	Username string `json:"username_user"`
	Password string `json:"password_user"`
	Email    string `json:"email_user"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}
