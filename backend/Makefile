PG_URL := ${POSTGRES_URL}

goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

goose-add:
	goose -dir ./migrations postgres "${PG_URL}" create hw-migration sql

goose-up:
	goose -dir ./migrations postgres "${PG_URL}" up

goose-down:
	goose -dir ./migrations postgres "${PG_URL}" down-to 0

goose-status:
	goose -dir ./migrations postgres "${PG_URL}" status

.PHONY: goose-install goose-add goose-up goose-down goose-status