package menu_console

import (
	//"fmt"
	//"os"
	//"path/filepath"

	//"fmt"
	//"fmt"
	"github.com/reeflective/console"
	"runtime"

	//"os"
	//"path/filepath"

	//"os"
	//"path/filepath"
	"time"
)

// setupPrompt is a function which sets up the prompts for the main menu.
func MySetupPrompt(m *console.Menu) {
	p := m.Prompt()

	p.Primary = func() string {
		//prompt := "\x1b[33mexample\x1b[0m [main] in \x1b[34m%s\x1b[0m\n> "
		prompt := "\x1b[33mFuckC2\x1b[0m \x1b[0m> "
		//wd, _ := os.Getwd()
		//
		//dir, err := filepath.Rel(os.Getenv("HOME"), wd)
		//if err != nil {
		//	dir = filepath.Base(wd)
		//}
		//
		//return fmt.Sprintf(prompt, dir)
		return prompt
	}

	p.Secondary = func() string { return ">" }
	p.Right = func() string {
		return "\x1b[1;30m" + time.Now().Format("03:04:05.000") + "\x1b[0m"
	}

	p.Transient = func() string { return "\x1b[1;30m" + ">> " + "\x1b[0m" }
}


// setupPrompt is a function which sets up the prompts for the main menu.
func SetupPrompt(m *console.Menu) {
	p := m.Prompt()

	p.Primary = func() string {
		prompt := "\x1b[33mFuckC2\x1b[0m \x1b[0m> "
		prompt_windows := "FuckC2 > "
		//prompt := "\x1b[33mexample\x1b[0m [main] in \x1b[34m%s\x1b[0m\n> "
		//wd, _ := os.Getwd()
		//
		//dir, err := filepath.Rel(os.Getenv("HOME"), wd)
		//if err != nil {
		//	dir = filepath.Base(wd)
		//}

		//return fmt.Sprintf(prompt, dir)
		if runtime.GOOS !="windows"{
			return prompt
		}

		return prompt_windows
	}

	p.Secondary = func() string { return ">" }
	p.Right = func() string {

		right_str := "\x1b[1;30m" + time.Now().Format("03:04:05.000") + "\x1b[0m"

		right_str_windows := "[" + time.Now().Format("03:04:05.000") + "]"

		//eturn "\x1b[1;30m" + time.Now().Format("03:04:05.000") + "\x1b[0m"
		if runtime.GOOS !="windows"{
			return right_str
		}

		return right_str_windows
	}

	p.Transient = func() string {

		Transient_str := "\x1b[1;30m" + ">> " + "\x1b[0m"

		Transient_str_windows := ">>"


		if runtime.GOOS !="windows"{
			return Transient_str
		}

		return Transient_str_windows

		//return "\x1b[1;30m" + ">> " + "\x1b[0m"

	}
}
