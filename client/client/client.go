package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	GatewayURL = "http://localhost:5000"
)

type Client struct {
	baseURL    string
	authToken  string
	httpClient *http.Client
}

func NewAPIClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) SetAuthToken() {
	secret := []byte("my-secret")
	claims := jwt.RegisteredClaims{
		Subject:   "animesh",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // 1-hour expiration
		Issuer:    "api-gateway",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Generated Token:", signedToken)
	c.authToken = signedToken
}

func (c *Client) sendRequest(method, path string, payload interface{}) error {
	var body io.Reader
	if payload != nil {
		payloadBytes, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("error marshaling payload: %w", err)
		}
		body = bytes.NewBuffer(payloadBytes)
	}

	url := fmt.Sprintf("%s%s", c.baseURL, path)
	fmt.Println("url: ", url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.authToken)
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(responseBody))
	}

	fmt.Printf("Response from %s: %s\n", path, string(responseBody))
	return nil
}
