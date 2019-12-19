package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductName string `json:"product_name" gorm:"column:product_name"`
	SKU         string `json:"sku" gorm:"column:sku"`
	Qty         int32  `json:"qty" gorm:"column:qty"`
}
