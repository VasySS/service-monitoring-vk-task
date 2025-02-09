build:
	docker-compose --env-file .env -f docker-compose.yml build $(c)

up:
	docker-compose --env-file .env -f docker-compose.yml up -d $(c)

down:
	docker-compose --env-file .env -f docker-compose.yml down $(c)

start:
	docker-compose --env-file .env -f docker-compose.yml start $(c)

stop:
	docker-compose --env-file .env -f docker-compose.yml stop $(c)

.PHONY: build up down start stop  