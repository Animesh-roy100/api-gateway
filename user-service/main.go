package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user", getUser)

	r.Run(":5002")

	fmt.Println("User service is listening on port 5002")
}

func getUser(*gin.Context) {
	fmt.Println("User data")
}
