package init_setting

import (
	pb "SimpleC2RpcTest/protobuf"
	"google.golang.org/grpc"
	"log"
)

type ClientCore struct {
	GrpcClient pb.ClientSendServerServiceClient
	GrpcConn *grpc.ClientConn
}

func (c *ClientCore)InitGrpc(Address string) error {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	//defer conn.Close()

	c.GrpcConn = conn


	// 建立gRPC连接
	c.GrpcClient = pb.NewClientSendServerServiceClient(conn)
	return nil
}