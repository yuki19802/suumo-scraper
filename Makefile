run:
	@mkdir -p data
	go run ./cmd/main.go

test:
	go test ./internal/...

