package command

import (
	"SimpleC2RpcTest/client_console/init_setting"
	"fmt"
)


//列出implants信息
func ListImplantsInfo(clientCore *init_setting.ClientCore)  {
	fmt.Println("in list implant info")

	//SendCommandInfo(clientCore)
	ClientRequestHostInfo(clientCore)

}