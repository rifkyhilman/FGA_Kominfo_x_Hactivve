package models

import (
	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	item_id       uint
	item_code     int64
	description   string
	quantity      string
	order_id      uint
	customer_name string
}
