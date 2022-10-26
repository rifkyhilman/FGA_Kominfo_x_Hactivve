package config

import (
	"Assigment_S2/models"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open("mysql", "root:@/orders_by?charset=utf8&parseTime=True&loc=local")

	if err != nil {
		panic("faild to connect database !!")
	} else {
		fmt.Println("succes connect to database !!!")
	}

	DB.AutoMigrate(&models.Items{})
}
