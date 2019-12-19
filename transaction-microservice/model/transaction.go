package model

import (
	"github.com/jinzhu/gorm"
)

type Transaction struct {
	gorm.Model
	ProductID       uint32 `json:"product_id" gorm:"column:product_id"`
	TransactionCode string `json:"transaction_code" gorm:"column:transaction_code"`
	Amount          int32  `json:"amount" gorm:"column:amount"`
	Product
}
