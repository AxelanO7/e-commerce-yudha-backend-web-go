package model

import (
	"gorm.io/gorm"
)

// Cart struct
type Cart struct {
	gorm.Model
	SaleId    uint    `json:"sale_id"`
	Sale      Sale    `gorm:"foreignKey:SaleId" json:"sale"`
	ProductId uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductId" json:"product"`
}

// Carts struct
type Carts struct {
	Carts []Cart `json:"carts"`
}
