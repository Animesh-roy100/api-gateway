package routes

import (
	"api-gateway/internal/adapter/application/rest/handler"
	"api-gateway/internal/adapter/application/rest/middleware"
	"api-gateway/internal/adapter/infrastucture/cache"
	"api-gateway/internal/adapter/infrastucture/ratelimit"
	"api-gateway/internal/domain/service"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func SetupRoutes(rg *gin.RouterGroup) {
	rateLimiter := ratelimit.NewTokenBucketLimiter(rate.Limit(100), 10)

	cacheRepo, err := cache.NewRedisCache(cache.RedisCacheOptions{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	})
	if err != nil {
		log.Fatalf("Failed to initialize Redis cache: %v", err)
	}

	// Gateway service
	gatewayService := service.NewGatewayService(cacheRepo)
	proxyHandler := handler.NewProxyHandler(gatewayService)

	// IP whitelist instance
	ipWhitelist := middleware.NewIPWhiteList()
	ipWhitelist.AddIPs([]string{"ALL"}) // Allowed all IPs

	// Apply middleware to all routes
	rg.Use(ipWhitelist.Middleware())
	rg.Use(middleware.Authenticate())
	rg.Use(middleware.RateLimit(rateLimiter))
	rg.Use(middleware.CircuitBreaker())

	// Define routes
	SetupPaymentRoutes(proxyHandler, rg)
	SetupUserRoutes(proxyHandler, rg)
	SetupProductRoutes(proxyHandler, rg)
}
