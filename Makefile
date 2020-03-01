
all: build_client buid_port_domain_service build_docker_compose

build_proto:
	cd portpb && make

build_client:
	cd client_api && make build

buid_port_domain_service:
	cd port_domain_service && make build

build_docker_compose:
	docker-compose build client_api port_domain_service

test:
	cd client_api && make test
	cd port_domain_service && make test

.PHONY: all build_proto build_client buid_port_domain_service build_docker_compose
