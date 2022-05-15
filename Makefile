HAS_YQ=$(shell which yq)
HAS_MIGRATE=$(shell which migrate)
NS=session

.PHONY: up
up:	 ## docker-compose up -d
	docker-compose up -d

.PHONY: build
build: ## docker-compose build
	docker-compose build

.PHONY: down
down: ## docker-compose down
	docker-compose down

.PHONY: bash
bash: ## bash
	docker exec -it usa-kabu-go-postgres-1 /bin/bash

.PHONY: migrate
migrate: ## migrate
ifeq ($(HAS_MIGRATE),)
	brew install golang-migrate
else
	migrate -help
	@echo 'migrate create -ext sql -dir ./migrations [NAME]'
	@echo 'migrate -database 'postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable' -path ./migrations up'
	@echo 'migrate -database 'postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable' -path ./migrations down'
	@echo 'migrate -database 'postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable' -path ./migrations force [バージョン番号]'
endif

.PHONY: schema-stocks
schema-stocks: ## schema-stocks
	go run entgo.io/ent/cmd/ent init Stocks

# gen
.PHONY: gen
gen: ## gen
	go get github.com/99designs/gqlgen@v0.17.5 && cd ./server && go run github.com/99designs/gqlgen

## generate
.PHONY: generate
generate:	## generate
	go generate ./server/generate.go

## gql
.PHONY: gql
gql: ## gql
	go generate ./server/ent/...

# air api
.PHONY: air-%
air-%: ## air -c $(@:air-%=%)/cmd/api/.air.toml
	air -c $(@:air-%=%)/cmd/api/.air.toml

# task
.PHONY: task
task:
	cd task/cmd/cli && go run main.go
