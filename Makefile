PREFIX ?= $(HOME)/go/bin
DATABASE_URL ?= "postgresql://$(USER):@localhost:5432/postgres?sslmode=disable"

init:
	@go mod tidy 
	@cd frontend && yarn && cd ..

build:
	@go build -o ./bin/app ./cmd/app

backend-dev:
	@$(PREFIX)/air --build.cmd "go build -o ./bin/app ./cmd/app" --build.bin "./bin/app api" --build.exclude_dir "templates,bin,migrations"

frontend-dev:
	@cd frontend && yarn dev

migrate:
	@migrate -path migrations/ -database $(DATABASE_URL) -verbose up 

rollback:
	@migrate -path migrations/ -database $(DATABASE_URL) -verbose down 1

sync:
	@pg_dump $(DATABASE_URL) --schema-only --no-owner --no-comments --no-acl > docs/schema.sql

.PHONY: all build run
