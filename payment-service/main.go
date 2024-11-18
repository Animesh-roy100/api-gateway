package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", makePayment)

	r.Run(":5003")

	fmt.Println("Payment service is listening on port 5003")
}

func makePayment(c *gin.Context) {
	var paymentData map[string]interface{}
	if err := c.BindJSON(&paymentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data"})
		return
	}

	response := gin.H{
		"status":  "success",
		"amount":  paymentData["amount"],
		"userId":  paymentData["userId"],
		"product": paymentData["productId"],
		"message": "Payment processed successfully",
	}

	fmt.Println("payment data", response)

	c.JSON(http.StatusOK, response)
}
