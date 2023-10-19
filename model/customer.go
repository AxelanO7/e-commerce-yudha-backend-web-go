package model

import (
	"gorm.io/gorm"
)

// Customer struct
type Customer struct {
	gorm.Model
	Name    string `json:"name_customer"`
	Address string `json:"address_customer"`
	Phone   string `json:"phone_customer"`
	UserID  uint   `json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}

// Customers struct
type Customers struct {
	Customers []Customer `json:"customers"`
}
