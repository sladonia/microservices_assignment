SERVICE_NAME=port_domain_service
BIN=app
BUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(BIN) ./src/.

run:
	@ go run ./src/.

build:
	$(BUILD)

docker_build:
	$(BUILD)
	docker build -t $(SERVICE_NAME) .

fmt:
	go fmt ./src/...

vet:
	go vet ./src/...

dep:
	@ cd src
	go mod tidy

test:
	@ go test ./src/services/...
	@ go test ./src/controllers/...

.PHONY: run build docker_build fmt vet dep test
