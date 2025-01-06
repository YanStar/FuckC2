package main

import (
	pb "SimpleC2RpcTest/protobuf"
	"context"
	"google.golang.org/grpc"
	"io"
	//"io"
	"log"
	//"time"
	//"time"
)

// Address 连接地址
const Address string = ":8000"

var grpcClient pb.ClientSendServerServiceClient

func RegisterClient() {
	req := pb.ClientRegister{
		ClientId: "yxx2",
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	stream, err := grpcClient.ClientRegisterService(context.Background())
	//stream, err := grpcClient.ClientRegisterService(ctx)

	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	err = stream.Send(&req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	// 接收服务器的响应
	resp, err := stream.Recv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	log.Printf("Received from server: %s", resp.Result)

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server closed the stream")
			return
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return
		}
		log.Printf("Message from server: %s", resp.Result)
	}
	//
	//go func() {
	//	for {
	//		resp, err := stream.Recv()
	//		if err == io.EOF {
	//			log.Println("Server closed the stream")
	//			return
	//		}
	//		if err != nil {
	//			log.Printf("Error receiving message: %v", err)
	//			return
	//		}
	//		log.Printf("Message from server: %s", resp.Result)
	//	}
	//}()

	//for {
	//	// 接收服务端响应
	//	//resp, err := stream.Recv()
	//	resp, _ := stream.Recv()
	//
	//	if resp != nil{
	//		log.Printf("Received from server: %s", resp.Result)
	//	}
	//
	//	if err == io.EOF {
	//		log.Println("Server closed the stream")
	//		//return
	//	}
	//	if err != nil {
	//		log.Printf("Failed to receive message: %v", err)
	//	}

	//if resp == nil{
	//	time.Sleep(10 * time.Second)
	//	continue
	//}else {
	//	log.Printf("Received from server: %s", resp.Result)
	//}
	//
	//time.Sleep(3 * time.Second)

	//}

}

func SendCommandInfo() {
	req := pb.ClientSendCommandInfo{
		ClientId:    "yxx2",
		CommandInfo: "shutdown",
	}

	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.ClientSendCommandInfoService(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewClientSendServerServiceClient(conn)

	//stream, err := grpcClient.ClientRegisterService(context.Background())

	go RegisterClient()
	SendCommandInfo()
	//time.Sleep(100)

	select {}

}
