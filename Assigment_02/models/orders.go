package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Items struct {
	gorm.Model
	item_id       uint
	item_code     int64
	description   string
	Quantity      string
	order_id      uint
	Customer_name string
	Ordered_at    time.Time
}

// type itemInput struct {
// 	Nama_customer string
// 	Harga         int
// }
