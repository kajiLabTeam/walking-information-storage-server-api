-include .env

up:
	docker compose build && docker compose up -d && docker compose logs -f

logs:
	docker compose logs -f

reup:
	rm -rf ./docker/db/data && docker compose build && docker compose up -d && docker compose logs -f

logs:
	docker compose logs -f

db:
	docker exec -it $(DB_HOST) psql -U $(DB_USER) -d $(DB_NAME)