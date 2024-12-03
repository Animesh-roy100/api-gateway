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
		"/products": "http://localhost:5001",
		"/users":    "http://localhost:5002",
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
	cacheKey := method + ":" + path

	// Try to retrieve from cache
	if cacheResponse, found := g.cache.Get(ctx, cacheKey); found {
		return &models.ServiceResponse{
			StatusCode: http.StatusOK,
			Body:       cacheResponse,
			Headers:    map[string][]string{"Cache": {"Hit"}},
		}, nil
	}

	response, err := g.MakeRequest(ctx, path, method, headers, body)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	expiration := 5 * time.Minute
	if err := g.cache.Set(ctx, cacheKey, response.Body, expiration); err != nil {
		return nil, fmt.Errorf("error setting cache: %w", err)
	}

	return response, nil
}

func (g *gatewayService) MakeRequest(ctx context.Context, path, method string, headers map[string][]string, body []byte) (*models.ServiceResponse, error) {
	targetService := ""
	for prefix, url := range g.serviceMap {
		if strings.HasPrefix(path, prefix) {
			targetService = url
			path = path[len(prefix):]
			break
		}
	}
	if targetService == "" {
		return nil, fmt.Errorf("no service found for path: %s", path)
	}

	fullURL := targetService + path
	fmt.Println("full url: ", fullURL)

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
