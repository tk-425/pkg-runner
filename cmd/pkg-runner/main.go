package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/tk-425/pkg-runner/internal/discover"
	"github.com/tk-425/pkg-runner/internal/tui"
)

var version = "dev"

func main() {
	ver := flag.Bool("v", false, "print version and exit")
	versionFlag := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			if v := info.Main.Version; v != "" && v != "(devel)" {
				version = v
			}
		}
	}

	if *ver || *versionFlag {
		fmt.Printf("pkg-runner %s\n", version)
		return
	}

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
