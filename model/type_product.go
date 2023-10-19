package model

import (
	"gorm.io/gorm"
)

// TypeProduct struct
type TypeProduct struct {
	gorm.Model
	Name string `json:"name_product_type"`
}

// TypeProducts struct
type TypeProducts struct {
	TypeProducts []TypeProduct `json:"product_types"`
}
