init:
	go mod tidy
	go mod vendor
	go run github.com/google/wire/cmd/wire@latest gen ./internal/nuk