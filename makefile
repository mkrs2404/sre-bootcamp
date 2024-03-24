.DEFAULT_GOAL := help

help: ## Display help
# This help commands picks up all the comments that start with `##` and prints them in a nice format.
# The comments should be in the following format:
# <target>:<space><comment>
# Taken from: https://stackoverflow.com/a/64996042/4257791
	@echo "Usage: make <target>"
	@echo "Targets:"
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

dc-up: ## dc-up starts the dependencies in the background
	docker-compose up -d

dc-down: ## dc-down stops the running dependencies
	docker-compose down

build: ## build builds the server binary
	go build -o server cmd/server/main.go

run: build ## run starts the server
	./server

check-lint: ## check-lint checks whether golangci-lint is installed
	@which golangci-lint || echo "Install golangci-lint from https://golangci-lint.run/usage/install/#local-installation"

lint: ## lint runs the linter
	golangci-lint run ./...

build-migration: ## build-migration builds the migration binary
	go build -o migration cmd/migration/main.go

migrate: build-migration ## migrate runs the up migration
	./migration

migrate-down: build-migration ## migrate-down runs the down migration. You can optionally pass the number of steps to rollback like: make migrate-down steps=1
	@if [ -z "$(steps)" ]; then ./migration --rollback; else ./migration --rollback --steps=$(steps); fi

force-migrate: build-migration ## force-migrate force migrates a schema version. It requires a version to be passed like: make force-migrate version=1
	FORCE_VERSION=$(version) ./migration


