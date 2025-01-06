package grpc_api

import (
	pb "SimpleC2RpcTest/protobuf"
	"fmt"
	"io"
	"log"
	"sync"
	"context"
)

// SimpleService 定义我们的服务
type ClientService struct {
	pb.UnimplementedClientSendServerServiceServer
	Command_info_chan  chan *pb.ClientSendCommandInfo
	Clients_stream_map map[string]pb.ClientSendServerService_ClientRegisterServiceServer
	mu                 sync.Mutex
}


// client 发来指令
func (s *ClientService) ClientSendCommandToImplantService(ctx context.Context, req *pb.ClientSendCommandInfo) (*pb.ClientRecvResultInfo, error) {


	fmt.Println("in server ClientSendCommandInfoService")

	//command_info_tmp := &pb.ClientSendCommandInfo{}

	s.Command_info_chan <- req

	res := pb.ClientRecvResultInfo{
		Result: "success you",
	}

	//close(s.command_info_chan)

	return &res, nil
}

//client 注册
func (s *ClientService) ClientRegisterService(stream pb.ClientSendServerService_ClientRegisterServiceServer) error {
	req, err := stream.Recv()
	if err == io.EOF {
		// 客户端关闭了流
		log.Println("Client disconnected")
		return nil
	}
	if err != nil {
		log.Printf("Error receiving from client: %v", err)
		return err
	}

	// 打印接收到的客户端信息
	log.Printf("Received from client %s register", req.ClientId)

	client_id := req.ClientId
	s.mu.Lock()
	s.Clients_stream_map[client_id] = stream // 将客户端 ID 和流关联
	s.mu.Unlock()
	log.Printf("Client %s connected save to map", client_id)

	//// 监听流上下文，清理失效连接
	//go func() {
	//	<-stream.Context().Done()
	//	log.Printf("Client '%s' disconnected", client_id)
	//	s.mu.Lock()
	//	delete(s.clients_stream_map, client_id)
	//	s.mu.Unlock()
	//}()

	res := pb.ClientRecvResultInfo{
		Result: "success you register",
	}

	// 发送响应到客户端
	if err := stream.Send(&res); err != nil {
		log.Printf("Error sending message to client: %v", err)
		return err
	}

	// 等待客户端消息
	for {
		// 这里可以监听客户端的后续请求
		_, err := stream.Recv()
		if err == io.EOF {
			// 客户端关闭了流
			break
		}
		if err != nil {
			log.Printf("Error receiving from client: %v", err)
			break
		}
	}
	// 客户端断开，清理连接
	s.mu.Lock()
	delete(s.Clients_stream_map, client_id)
	s.mu.Unlock()
	log.Printf("Client %s disconnected", client_id)
	return nil

}

//client 查询hostinfo
func (s *ClientService) ClientRequestHostInfo(ctx context.Context, req *pb.ClientRequestSingleCommand) (*pb.HostInfoListResponse, error) {

	fmt.Println("in server ClientRequestHostInfo")

	// 模拟一些 HostList 数据
	hosts := []*pb.HostInfo{
		{
			ClientId:  "3",
			Hostname:  "implant_2",
			Ip:        "127.0.0.1",
			ConnPort:  "8888",
			Os:        "windows",
			Privilege: "admin",
			Version:   "0.1",
			Remarks:   "shuyang",
		},
		{
			ClientId:  "4",
			Hostname:  "implant_1",
			Ip:        "127.0.0.1",
			ConnPort:  "8888",
			Os:        "windows",
			Privilege: "admin",
			Version:   "0.1",
			Remarks:   "shuyang",
		},
	}

	res := pb.HostInfoListResponse{
		HostinfoList: hosts,
	}

	//close(s.command_info_chan)

	return &res, nil
}