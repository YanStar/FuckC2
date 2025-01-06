package menu_console

import (
	"bufio"
	"fmt"
	"github.com/reeflective/console"
	"os"
	"strings"
)

func ExitCtrlD(c *console.Console) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Confirm exit (Y/y): ")
	text, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(text)

	if (answer == "Y") || (answer == "y") {
		os.Exit(0)
	}
}