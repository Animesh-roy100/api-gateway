package main

import (
	"client/client"
	"log"
)

// Created client to test the api-gateway
func main() {
	apiClient := client.NewAPIClient(client.GatewayURL)
	apiClient.SetAuthToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGktZ2F0ZXdheSIsInN1YiI6InVzZXIxMjMiLCJleHAiOjE3MzE5NjMyMDJ9.EkJB4IHkAuoVB5O2KTmxeK0vnQ3mZeOl_J39ocfOEXM")

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
