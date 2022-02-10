package models

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name          string
	Quantity      int8
	Selling_Price string
}
