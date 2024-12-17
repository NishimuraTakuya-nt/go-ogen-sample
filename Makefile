.PHONY: go-download oas-gen help

# 環境変数
NAME := go-ogen-sample
DC := docker compose


# =================== CI ===================
go-download: ## Download dependencies
	go mod download
	go mod verify
	go mod tidy


# =================== generate ===================
oas-gen: ## Generate openapi
	go generate ./...


# =================== container ===================
docker-down: ## Stop container
	$(DC) down

docker-build: ## Build image
	$(DC) build $(NAME)

docker-up: ## Run container
	$(DC) up -d



help: ## display this help.
    @awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
