.PHONY: go-download help

# =================== CI ===================
go-download: ## Download dependencies
	go mod download
	go mod verify
	go mod tidy

# =================== generate ===================
oas-gen: ## Generate openapi
	go generate ./...


help: ## display this help.
    @awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
