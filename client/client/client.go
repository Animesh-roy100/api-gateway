package client

import (
	"net/http"
	"time"
)

const (
	GatewayURL = "http://localhost:5000/api/v1"
	AuthToken  = "auihefrwiqoleabfpqwlr3bgvhjavrwiqrd"
)

type Client struct {
	baseURL    string
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
