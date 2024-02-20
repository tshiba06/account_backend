-include .env
export

ARGS ?= default

.PHONEY: build
build:
	docker compose up -d --build

.PHONEY: up
up:
	docker compose up -d

.PHONEY: test
test:
	go test ./... -count=1

# TODO: 回数指定ができるとよい、defaultはallで
.PHONEY: migrate-up
migration-up:
	docker run -u $(id -u):$(id -g) -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB) up

.PHONEY: create-migrate
create-migrate:
	docker run -u $(id -u):$(id -g) -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations/ create -ext sql -dir ./migrations -seq $(ARGS)

# WSLの場合にはdockerで動かすとなぜか-u指定していてもroot権限でファイルが作成されるので一旦簡易的にこのコマンドで処理をする
.PHONEY: chown
chown:
	sudo chown -R $(id -u):$(id -g) ARGS
