// internal/domain/service/gateway_config.go
package service

import "time"

type ServiceConfig struct {
	BaseURL     string
	Timeout     time.Duration
	RateLimit   int
	RequireAuth bool
}

type GatewayConfig struct {
	Services map[string]ServiceConfig
	Cache    struct {
		Enabled    bool
		DefaultTTL time.Duration
	}
	CircuitBreaker struct {
		Enabled          bool
		FailureThreshold float64
		ResetTimeout     time.Duration
	}
}

func NewDefaultConfig() *GatewayConfig {
	return &GatewayConfig{
		Services: map[string]ServiceConfig{
			"/users": {
				BaseURL:     "http://localhost:5002",
				Timeout:     time.Second * 30,
				RateLimit:   100,
				RequireAuth: true,
			},
			"/products": {
				BaseURL:     "http://localhost:5001",
				Timeout:     time.Second * 30,
				RateLimit:   200,
				RequireAuth: true,
			},
			"/payments": {
				BaseURL:     "http://localhost:5003",
				Timeout:     time.Second * 30,
				RateLimit:   50,
				RequireAuth: true,
			},
		},
		Cache: struct {
			Enabled    bool
			DefaultTTL time.Duration
		}{
			Enabled:    true,
			DefaultTTL: time.Minute * 5,
		},
		CircuitBreaker: struct {
			Enabled          bool
			FailureThreshold float64
			ResetTimeout     time.Duration
		}{
			Enabled:          true,
			FailureThreshold: 0.5,
			ResetTimeout:     time.Second * 60,
		},
	}
}
