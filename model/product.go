package model

import (
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Name          string      `json:"name_product"`
	Amount        int         `json:"product_amount"`
	TypeProductID uint        `json:"product_type_id"`
	TypeProduct   TypeProduct `gorm:"foreignKey:TypeProductID"`
	Price         int         `json:"product_price"`
}

// Products struct
type Products struct {
	Products []Product `json:"products"`
}
