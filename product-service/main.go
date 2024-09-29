package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/product", getProduct)

	r.Run(":5001")

	fmt.Println("Product service is listening on port 5001")
}

func getProduct(*gin.Context) {
	fmt.Println("Product data")
}
