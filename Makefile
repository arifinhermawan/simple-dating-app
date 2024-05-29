.PHONY: build

# for dependencies
dep:
	@echo "RUNNING GO MOD..."
	@go mod tidy
	@go mod vendor

run:
	@go run cmd\main.go