package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:id", getProduct)

	r.Run(":5001")

	fmt.Println("Product service is listening on port 5001")
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	response := gin.H{
		"id":          id,
		"name":        "Animesh",
		"description": "good product",
		"price":       100,
		"currency":    "INR",
		"message":     "Product details fetched successfully",
	}
	fmt.Println("Product data: ", response)

	c.JSON(http.StatusOK, response)
}
