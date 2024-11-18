package main

import "log"

const (
	gatewayURL = "http://localhost:5000/api/v1"
	authToken  = "auihefrwiqoleabfpqwlr3bgvhjavrwiqrd"
)

func main() {
	client := NewAPIClient(gatewayURL)

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
