package internal

import (
	"gorm.io/gorm"
)

type Product struct {
    gorm.Model
    Item     string   `json:"item" gorm:"not null;unique"`
    Price    float32  `json:"price" gorm:"not null"`
    Quantity int      `json:"quantity" gorm:"not null"`
    Supplier Supplier `gorm:"foreignKey:ProductID" json:"supplier"`
}

type Supplier struct {
    Name      string `json:"name" gorm:"not null"`
    Address   string `json:"address" gorm:"not null"`
    Tel       string `json:"tel" gorm:"not null"`
    ProductID uint
}

