package main

import (
	"project-pertama/Assigment_02/config"
	"project-pertama/Assigment_02/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectionDB()

	router.GET("/", controllers.GetItems)
	router.GET("/order/:id", controllers.GetItemsbyID)
	router.POST("/orders", controllers.PostItems)

	router.Run()
}
