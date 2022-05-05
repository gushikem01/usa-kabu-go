HAS_YQ=$(shell which yq)
HAS_MIGRATE=$(shell which migrate)
NS=session

.PHONY: up
up:	 ## docker-compose up -d
	docker-compose up -d

.PHONY: build
build:	 ## docker-compose build
	docker-compose build

.PHONY: down
down:	 ## docker-compose down
	docker-compose down

.PHONY: bash
bash:	## bash
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

.PHONY: generate
generate:	## generate
	go generate ./ent

# go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
