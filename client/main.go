package main

import (
	"client/client"
	"log"
)

// Created client to test the api-gateway
func main() {
	apiClient := client.NewAPIClient(client.GatewayURL)
	apiClient.SetAuthToken()

	// user service
	if err := apiClient.GetUser("1"); err != nil {
		log.Printf("Error getting user: %v", err)
	}

	// product service
	if err := apiClient.GetProduct("5"); err != nil {
		log.Printf("Error getting product: %v", err)
	}

	// payment service
	payment := map[string]interface{}{
		"amount":      100,
		"currency":    "INR",
		"productId":   "5",
		"userId":      "1",
		"paymentType": "cash",
	}
	if err := apiClient.MakePayment(payment); err != nil {
		log.Printf("Error making payment: %v", err)
	}
}
