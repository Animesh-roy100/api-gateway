package main

import (
	"client/client"
	"log"
)

func main() {
	client := client.NewAPIClient(client.GatewayURL)

	// user service
	if err := client.GetUser("1"); err != nil {
		log.Printf("Error getting user: %v", err)
	}

	// product service
	if err := client.GetProduct("5"); err != nil {
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
	if err := client.MakePayment(payment); err != nil {
		log.Printf("Error making payment: %v", err)
	}
}
