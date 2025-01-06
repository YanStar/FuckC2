package command

import (
	"github.com/reeflective/console"
	"github.com/reeflective/console/commands/readline"

	//"github.com/reeflective/console/commands/readline"
	"github.com/spf13/cobra"
	"SimpleC2RpcTest/client_console/menu_console"
	"SimpleC2RpcTest/client_console/init_setting"

)

var shortUsage = "FuckC2 console test"

func MainMenuCommands(app *console.Console, clientCore *init_setting.ClientCore) console.Commands {

	return func() *cobra.Command {

		rootCmd := &cobra.Command{}

		rootCmd.Short = shortUsage


		rootCmd.AddGroup(
			&cobra.Group{ID: "Generic", Title: "Generic"},
			&cobra.Group{ID: "Commands", Title: "Commands"},
			//&cobra.Group{ID: "filesystem", Title: "filesystem"},
			//&cobra.Group{ID: "deployment", Title: "deployment"},
			//&cobra.Group{ID: "tools", Title: "tools"},
		)

		rootCmd.AddCommand(readline.Commands(app.Shell()))


		exitCmd := &cobra.Command{
			Use:     "exit",
			Short:   "Exit the console application",
			GroupID: "Commands",
			Run: func(cmd *cobra.Command, args []string) {
				menu_console.ExitCtrlD(app)
			},
		}
		rootCmd.AddCommand(exitCmd)


		listCmd := &cobra.Command{
			Use:     "list",
			Short:   "list implants info",
			GroupID: "Generic",
			Run: func(cmd *cobra.Command, args []string) {
				//menu_console.ExitCtrlD(app)
				ListImplantsInfo(clientCore)

			},
		}
		rootCmd.AddCommand(listCmd)










		return rootCmd

	}



}