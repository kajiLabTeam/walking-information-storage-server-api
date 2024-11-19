-include .env

up:
	docker compose build && docker compose up -d

reup:
	rm -rf ./docker/postgres/data && docker compose build && docker compose up -d

logs:
	docker compose logs -f

db:
	docker exec -it $(DB_HOST) psql -U $(DB_USER) -d $(DB_NAME)

go:
	docker exec -it $(GO_HOST) /bin/sh