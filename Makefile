-include .env
export

ifndef ARGS
$(error ARGS is not set)
endif

.PHONEY: build
build:
	docker compose up -d --build

.PHONEY: test
test:
	docker compose exec -ti server go test ./... -count=1

.PHONEY: up
up:
	docker compose up -d

# TODO: 回数指定ができるとよい、defaultはallで
.PHONEY: migrate-up
migration-up:
	docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB) up

.PHONEY: create-migrate
create-migrate:
	docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ create -ext sql -dir ./migrations -seq $(ARGS)
