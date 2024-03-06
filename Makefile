.PHONY: init dev

init:
	go mod tidy
	go mod vendor
	go run github.com/google/wire/cmd/wire@latest gen ./internal
	if [ ! -f .env ]; then cp .env.example .env; fi
dev:
	go run github.com/cosmtrek/air@latest -c .air.toml