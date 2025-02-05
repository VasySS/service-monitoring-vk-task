build:
	docker-compose -f docker-compose.yml build $(c)

up:
	docker-compose -f docker-compose.yml up -d $(c)

down:
	docker-compose -f docker-compose.yml down $(c)

start:
	docker-compose -f docker-compose.yml start $(c)

stop:
	docker-compose -f docker-compose.yml stop $(c)

.PHONY: build up down start stop  