package model

import (
	"gorm.io/gorm"
)

// Sale struct
type Sale struct {
	gorm.Model
	SaleAmount      int      `json:"sale_amount"`
	DateTransaction string   `json:"date_transaction"`
	UserID          uint     `json:"user_id"`
	User            User     `gorm:"foreignKey:UserID" json:"user"`
	ProductID       uint     `json:"product_id"`
	Product         Product  `gorm:"foreignKey:ProductID" json:"product"`
	CustomerID      uint     `json:"customer_id"`
	Customer        Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Price           int      `json:"price"`
	TotalSale       int      `json:"total_sale"`
}

// Sales struct
type Sales struct {
	Sales []Sale `json:"sales"`
}
