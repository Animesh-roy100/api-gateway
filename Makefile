# Define directories
CLIENT_DIR := ./client
API_GATEWAY_DIR := ./api-gateway
USER_SERVICE_DIR := ./user-service
PRODUCT_SERVICE_DIR := ./product-service
PAYMENT_SERVICE_DIR := ./payment-service

# Default port assignments
API_GATEWAY_PORT := 5000
USER_SERVICE_PORT := 5002
PRODUCT_SERVICE_PORT := 5001
PAYMENT_SERVICE_PORT := 5003

# Targets
.PHONY: all client api-gateway services user-service product-service payment-service clean

# Run all components
all: services api-gateway client

# Run the client
client:
	@echo "Starting Client..."
	cd $(CLIENT_DIR) && go run main.go

# Run the API Gateway
api-gateway:
	@echo "Starting API Gateway on port $(API_GATEWAY_PORT)..."
	cd $(API_GATEWAY_DIR) && go run cmd/main.go

# Run all services
services: user-service product-service payment-service

# Run the User Service
user-service:
	@echo "Starting User Service on port $(USER_SERVICE_PORT)..."
	cd $(USER_SERVICE_DIR) && go run main.go

# Run the Product Service
product-service:
	@echo "Starting Product Service on port $(PRODUCT_SERVICE_PORT)..."
	cd $(PRODUCT_SERVICE_DIR) && go run main.go

# Run the Payment Service
payment-service:
	@echo "Starting Payment Service on port $(PAYMENT_SERVICE_PORT)..."
	cd $(PAYMENT_SERVICE_DIR) && go run main.go

# Clean (optional, in case you have build artifacts)
clean:
	@echo "Cleaning up build artifacts..."
	find . -name "*.out" -type f -delete
	find . -name "*.exe" -type f -delete
