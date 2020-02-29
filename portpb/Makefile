all:
	protoc --go_out=plugins=:. port.proto
	cp -rp port.pb.go ../port_domain_service/src/portpb/port.pb.go
	cp -rp port.pb.go ../client_api/src/portpb/port.pb.go

.PHONY: all
