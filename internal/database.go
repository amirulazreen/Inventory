package internal

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Item     string   `json:"item"`
	Price    float32  `json:"price"`
	Quantity int      `json:"quantity"`
	Supplier Supplier `gorm:"foreignKey:ProductID" json:"supplier"`
}

type Supplier struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Tel       string `json:"tel"`
	ProductID uint
}