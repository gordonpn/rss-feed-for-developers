.PHONY: up-pg

up-pg:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml run --rm --publish=127.0.0.1:5432:5432 postgres

clean-db:
	docker volume rm dev-postgres-rss-feed
	docker volume create --name=dev-postgres-rss-feed
