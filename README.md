# microservices_assignment

### api usage

import ports
```sh
curl -X POST -d @ports.json localhost:8080/ports

# example response:
{
    "number_inserted":1632,
    "number_updated":0
}
```

retrieve port
```sh
curl localhost:8080/port/AEAJM

# example response
{
    "abbreviation": "AEAJM",
    "name": "Ajman",
    "coordinates": [
        54.37,
        24.47
    ],
    "city": "Ajman",
    "province": "Ajman",
    "country": "United Arab Emirates",
    "timezone": "Asia/Dubai",
    "unlocs": [
        "AEAUH"
    ]
}
```