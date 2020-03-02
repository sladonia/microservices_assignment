package grpc_client

import (
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

// InitGrpcConnection inits grpc connection
func InitGrpcConnection(host, port string) (err error) {
	// TODO: implement connection healthcheck and reconnection
	Conn, err = grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	return err
}
