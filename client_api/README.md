# client_api

REST api service to import and retrieve port data

### endpoints

import ports (expects json body of known format):
```sh
POST /ports
```

retrieve ports:
```sh
GET /port/{port_abbreviation}
```

### service dependencies

depends on port_domain_service

### build

build executable & docker container
```sh
make docker_build
```

### testing

run tests
```sh
make test
```

### configuration

service configured with the env vars and the defaults are:
```sh
SERVICE_NAME=client_api
ENV=dev
LOG_LEVEL=debug
PORT=:8080
SHUTDOWN_TIMEOUT=20 # greacefull shutdown timeout
PORT_DOMAIN_HOST=localhost
PORT_DOMAIN_PORT=50051
```