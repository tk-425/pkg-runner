package runner

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/tk-425/pkg-runner/internal/discover"
)

type Runner struct {
	stdout io.Writer
	stderr io.Writer
}

func New() *Runner {
	return &Runner{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
}

func (r *Runner) SetOutput(w io.Writer) {
	r.stdout = w
	r.stderr = w
}

func (r *Runner) Run(ctx context.Context, script discover.Script) (int, error) {
	args := strings.Fields(script.Command)
	if len(args) == 0 {
		return 1, nil
	}
	// #nosec G204 — commands originate from trusted local manifest files
	// nosemgrep: dangerous-exec-command
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdout = r.stdout
	cmd.Stderr = r.stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode(), nil
		}
		if ctx.Err() != nil {
			return -1, nil
		}
		return 1, err
	}

	return 0, nil
}
