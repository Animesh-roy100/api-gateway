package service

import (
	"api-gateway/internal/domain/models"
	"api-gateway/internal/domain/port"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
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
	// Check cache first
	if method == "GET" {
		if cachedResponse, exists := g.cache.Get(ctx, path); exists {
			return &models.ServiceResponse{
				StatusCode: http.StatusOK,
				Body:       cachedResponse,
				Headers:    map[string][]string{"Content-Type": {"application/json"}},
			}, nil
		}
	}

	// Find target service
	targetService := ""
	for prefix, url := range g.serviceMap {
		if pathHasPrefix(path, prefix) {
			targetService = url
			break
		}
	}

	if targetService == "" {
		return nil, fmt.Errorf("no service found for path: %s", path)
	}

	// Create new request
	req, err := http.NewRequestWithContext(
		ctx,
		method,
		targetService+path,
		bytes.NewReader(body),
	)
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

	// Cache GET responses
	if method == "GET" && resp.StatusCode == http.StatusOK {
		g.cache.Set(ctx, path, responseBody, time.Minute*5) // Cache for 5 minutes
	}

	return &models.ServiceResponse{
		StatusCode: resp.StatusCode,
		Body:       responseBody,
		Headers:    resp.Header,
	}, nil
}

func pathHasPrefix(path, prefix string) bool {
	if len(path) < len(prefix) {
		return false
	}
	if len(path) == len(prefix) {
		return path == prefix
	}
	return path[0:len(prefix)] == prefix && (path[len(prefix)] == '/' || prefix == "/")
}
