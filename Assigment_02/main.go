package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", GetItems)
	router.GET("/order/:id", GetItemsbyID)
	router.POST("/orders", PostItems)

	router.Run()
}

func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Nama": "Gus Samsyudin",
	})
}

func GetItemsbyID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type itemInput struct {
	Nama_customer string
	Harga         int
}

func PostItems(c *gin.Context) {
	var itemInput itemInput

	err := c.ShouldBindJSON(&itemInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Nama_customer": itemInput.Nama_customer,
		"harga":         itemInput.Harga,
	})
}
