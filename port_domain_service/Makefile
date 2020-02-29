SERVICE_NAME=port_domain_service
BIN=app
BUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(BIN) ./src/.

run:
	@ go run ./src/.

build:
	$(BUILD)

fmt:
	go fmt ./src/...

dep:
	@ cd src
	go mod tidy

test:
	@ go test ./src/...

.PHONY: run build fmt dep test
