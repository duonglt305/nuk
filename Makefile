.PHONY: init dev migration

init:
	go mod tidy && go mod vendor
	go run github.com/google/wire/cmd/wire@latest gen ./internal
	if [ ! -f .env ]; then cp .env.example .env; fi
dev:
	@go run github.com/air-verse/air@latest -c .air.toml

migration:
	@go run github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -ext sql -dir db/migrations $(name)
migrate:
	@go run cmd/migrate/main.go $(t)
