.PHONY: up-pg clean-db up pg help
.DEFAULT_GOAL := help

up: ## Start the docker compose stack for development
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

up-pg: ## Only start Postgres for development
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml run --rm --publish=127.0.0.1:5432:5432 postgres

clean-db: ## Clean database
	docker volume rm dev-postgres-rss-feed
	docker volume create --name=dev-postgres-rss-feed

pg: ## Postgres interactive terminal
	docker exec -it rss-feed-postgres-dev sh

help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[33m%-20s\033[0m %s\n", $$1, $$2}'
