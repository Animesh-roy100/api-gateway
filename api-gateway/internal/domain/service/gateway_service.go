package service

import (
	"api-gateway/internal/domain/models"
	"api-gateway/internal/domain/port"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type gatewayService struct {
	cache      port.CacheRepository
	serviceMap map[string]string
	httpClient *http.Client
}

func NewGatewayService(cache port.CacheRepository) GatewayService {
	// Initialize service routes
	serviceMap := map[string]string{
		"/users":    "http://localhost:5002",
		"/products": "http://localhost:5001",
		"/payments": "http://localhost:5003",
	}

	return &gatewayService{
		cache:      cache,
		serviceMap: serviceMap,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (g *gatewayService) ValidateRequest(ctx context.Context, path string, method string) error {
	// TODO: Implement validate request
	return nil
}

func (g *gatewayService) ProxyRequest(ctx context.Context, path string, method string, headers map[string][]string, body []byte) (*models.ServiceResponse, error) {
	// Find target service
	var targetService string
	normalizedPath := path

	for prefix, url := range g.serviceMap {
		if strings.HasPrefix(path, prefix) {
			targetService = url
			normalizedPath = path[len(prefix):]
			break
		}
	}

	if targetService == "" {
		return nil, fmt.Errorf("no service found for path: %s", path)
	}

	// Construct the full URL
	fullURL := targetService + normalizedPath

	// Create new request
	req, err := http.NewRequestWithContext(ctx, method, fullURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Copy headers
	for key, values := range headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Send request
	resp, err := g.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return &models.ServiceResponse{
		StatusCode: resp.StatusCode,
		Body:       responseBody,
		Headers:    resp.Header,
	}, nil
}
