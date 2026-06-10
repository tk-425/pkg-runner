package main

import (
	"fmt"
	"os"

	"github.com/tk-425/pkg-runner/internal/discover"
	"github.com/tk-425/pkg-runner/internal/tui"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: unable to get working directory:", err)
		os.Exit(1)
	}

	scripts, err := discover.Discover(cwd)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}

	if len(scripts) == 0 {
		if err != nil {
			os.Exit(1)
		}
		fmt.Println("No scripts found")
		return
	}

	p := tui.NewProgram(scripts)
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
