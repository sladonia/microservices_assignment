# microservices_assignment

test assignment


### build dependencies

* go v1.13
* docker, docker-compose
* GNU make

### build

to build services and docker containers run:
```sh
make
```

### up services

run:
```sh
docker-compose up -d client_api
```

### api usage

import ports
```sh
curl -X POST -d @ports.json localhost:8080/ports

# example response:
{
    "number_inserted":1632,
    "number_updated":0,
    "encounter_errors":false
}
```

retrieve port
```sh
curl localhost:8080/port/AEAJM

# example response
{
    "abbreviation": "AEAJM",
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
        54.37,
        24.47
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
        "AEAUH"
    ],
    "code": "52000"
}
```

### testing

please note, port_domain_service require mongo_test service running to run tests:
```sh
docker-compose up -d mongo_test
``` 

run tests:
```sh
make test
```

### disclaimer

The project and it's structure apply a number of simplifications. It's only purpose is to demonstrate programming skills. For the real word scenario:
* packages client_api, port_domain_service, portpb should be separate repositories
* portpb package as a common dependency should be installed using `go get`

### further improvements

* implement grpc message streaming in chunks and benchmark the entire system to find the best performance option 
* implement grpc connection healthcheck and reconnect
* implement request_id generation middleware
