.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD TIDY..."
	@go mod tidy

	@echo "RUNNING GO MOD VENDOR..."
	@go mod vendor

run:
	@echo "Starting Docker containers...";\
	docker-compose up -d; 

	@echo "Running Go application..."
	@go run main.go
