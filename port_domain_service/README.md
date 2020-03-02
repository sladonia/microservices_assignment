# port_domain_service

grpc domain service

### service dependencies

depends on mongodb

### build

build executable & docker container
```sh
make docker_build
```

### testing

please note, port_domain_service require mongo_test service running to run tests:
```sh
docker-compose up -d mongo_test
``` 

run tests
```sh
make test
```

### configuration

service configured with the env vars and the defaults are:
```sh
SERVICE_NAME=port_domain_service
ENV=dev
LOG_LEVEL=debug
PORT=:50051
MONGO_USER=user
MONGO_PASSWORD=password
MONGO_HOST=localhost
MONGO_PORT=27017
MONGO_DB_NAME=port_db
```