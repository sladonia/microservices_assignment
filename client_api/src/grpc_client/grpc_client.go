package grpc_client

import (
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn

func InitGrpcClient(host, port string) (err error) {
	Conn, err = grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	return err
}
