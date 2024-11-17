package service

import (
	"api-gateway/internal/domain/models"
	"context"
)

type GatewayService interface {
	ProxyRequest(ctx context.Context, path string, method string, headers map[string][]string, body []byte) (*models.ServiceResponse, error)
	ValidateRequest(ctx context.Context, path string, method string) error
}
