package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) MakePayment(payment map[string]interface{}) error {
	payloadBytes, err := json.Marshal(payment)
	if err != nil {
		return fmt.Errorf("error marshaling payment: %w", err)
	}

	url := fmt.Sprintf("%s/payments", gatewayURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	fmt.Printf("Payment Response: %s\n", string(body))
	return nil
}
