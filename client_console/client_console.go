package main

import (
	//pb "SimpleC2RpcTest/protobuf"
	"fmt"
	"github.com/reeflective/console"
	"SimpleC2RpcTest/client_console/menu_console"
	"io"
	"SimpleC2RpcTest/client_console/command"
	"SimpleC2RpcTest/client_console/init_setting"
)


var clientCore *init_setting.ClientCore // 包级别的变量

func init(){
	fmt.Println("i am init")

	clientCore = &init_setting.ClientCore{}
	clientCore.InitGrpc(":8000")

}


func main()  {
	fmt.Println("hello console")

	// 在 main 中使用 clientCore
	if clientCore.GrpcClient != nil {
		fmt.Println("clientCore GrpcClient initialized successfully.")
	} else {
		fmt.Println("clientCore GrpcClient is nil.")
	}

	defer clientCore.GrpcConn.Close()


	app := console.New("FuckC2")

	// Global Setup ------------------------------------------------- //
	app.NewlineBefore = true
	app.NewlineAfter = true

	app.SetPrintLogo(func(_ *console.Console) {
		fmt.Print(`
 $$$$$$$$\                  $$\        $$$$$$\   $$$$$$\  
$$  _____|                 $$ |      $$  __$$\ $$  __$$\ 
$$ |   $$\   $$\  $$$$$$$\ $$ |  $$\ $$ /  \__|\__/  $$ |
$$$$$\ $$ |  $$ |$$  _____|$$ | $$  |$$ |       $$$$$$  |
$$  __|$$ |  $$ |$$ /      $$$$$$  / $$ |      $$  ____/ 
$$ |   $$ |  $$ |$$ |      $$  _$$<  $$ |  $$\ $$ |      
$$ |   \$$$$$$  |\$$$$$$$\ $$ | \$$\ \$$$$$$  |$$$$$$$$\ 
\__|    \______/  \_______|\__|  \__| \______/ \________|
                                                         

`)
	})


	// Main Menu Setup ---------------------------------------------- //

	menu := app.ActiveMenu()


	menu_console.MySetupPrompt(menu)

	hist, _ := menu_console.EmbeddedHistory(".example-history")
	menu.AddHistorySource("local history", hist)

	menu.AddInterrupt(io.EOF, menu_console.ExitCtrlD)



	menu.SetCommands(command.MainMenuCommands(app,clientCore))


	app.Start()

}
