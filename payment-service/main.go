package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/payment", getPayment)

	r.Run(":5003")

	fmt.Println("Payment service is listening on port 5003")
}

func getPayment(*gin.Context) {
	fmt.Println("Payment data")
}
