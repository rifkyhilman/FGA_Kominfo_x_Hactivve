package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("DB Connection error")
	}

	fmt.Println("Database Konek Brooo !!!")
}
