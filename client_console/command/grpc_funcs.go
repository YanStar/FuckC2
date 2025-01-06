package command

import (
	pb "SimpleC2RpcTest/protobuf"
	"SimpleC2RpcTest/client_console/init_setting"

	"log"
	"context"
	"SimpleC2RpcTest/client_console/common"
)

//client 发送命令给server
func SendCommandInfo(clientCore *init_setting.ClientCore) {
	req := pb.ClientSendCommandInfo{
		ClientId:    "yxx",
		CommandInfo: "shutdown",
	}

	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	//res, err := clientCore.grpcClient.ClientSendCommandInfoService(context.Background(), &req)
	res, err := clientCore.GrpcClient.ClientSendCommandToImplantService(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}


func ClientRequestHostInfo(clientCore *init_setting.ClientCore) {
	req := pb.ClientRequestSingleCommand{
		RequestCommand: "list",
	}

	res, err := clientCore.GrpcClient.ClientRequestHostInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}


	//// 打印返回值
	//log.Println(res)
	//
	//for _,host := range res.HostinfoList{
	//	log.Println(host)
	//
	//}



	common.PrintTable(res.HostinfoList,-1)

}