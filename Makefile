default:
	@echo See Makefile

fmt:
	@go fmt ./pkg/...
	@go fmt ./cmd/...

scrape:
	@mkdir -p data
	@go run ./cmd/scrape/main.go

import:
	@go run ./cmd/import/main.go

test:
	go test ./internal/...

################################################################################
# Elasticsearch helpers
es:
	@echo "Starting Elasticsearch stack"
	cd ./elastic && docker-compose up

clear-es:
	@echo "Deleting ES indices, some may not be found and that's fine"
	curl -XDELETE 'http://localhost:9200/listings'
	@echo

