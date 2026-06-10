package runner

import (
	"context"
	"io"
	"os"
	"os/exec"

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
	cmd := exec.CommandContext(ctx, "sh", "-c", script.Command)
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
