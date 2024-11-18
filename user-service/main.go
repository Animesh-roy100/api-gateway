package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:id", getUser)

	r.Run(":5002")

	fmt.Println("User service is listening on port 5002")
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	response := gin.H{
		"id":      id,
		"name":    "Animesh",
		"email":   "animesh@gmail.com",
		"role":    "customer",
		"message": "User details fetched successfully",
	}
	fmt.Println("User data: ", response)
	c.JSON(http.StatusOK, response)
}
