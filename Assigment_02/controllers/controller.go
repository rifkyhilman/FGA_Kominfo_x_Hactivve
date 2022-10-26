package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"project-pertama/Assigment_02/models"
)

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

func PostItems(c *gin.Context) {
	var ItemInput models.Items

	err := c.ShouldBindJSON(&ItemInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Nama_customer": ItemInput.Customer_name,
		"quantity":      ItemInput.Quantity,
	})
}
