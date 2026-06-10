package runner

import (
	"bytes"
	"context"
	"testing"

	"github.com/tk-425/pkg-runner/internal/discover"
)

func TestRunner_SuccessfulCommand(t *testing.T) {
	var out bytes.Buffer
	r := New()
	r.SetOutput(&out)

	script := discover.Script{
		Name:    "hello",
		Command: "echo hello",
		Source:  "test",
	}

	code, err := r.Run(context.Background(), script)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != 0 {
		t.Errorf("expected exit 0, got %d", code)
	}
	if out.String() != "hello\n" {
		t.Errorf("expected output %q, got %q", "hello\n", out.String())
	}
}

func TestRunner_FailingCommand(t *testing.T) {
	var out bytes.Buffer
	r := New()
	r.SetOutput(&out)

	script := discover.Script{
		Name:    "fail",
		Command: "exit 42",
		Source:  "test",
	}

	code, err := r.Run(context.Background(), script)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != 42 {
		t.Errorf("expected exit 42, got %d", code)
	}
}

func TestRunner_ContextCancellation(t *testing.T) {
	var out bytes.Buffer
	r := New()
	r.SetOutput(&out)

	script := discover.Script{
		Name:    "sleep",
		Command: "sleep 10",
		Source:  "test",
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	code, err := r.Run(ctx, script)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code == 0 {
		t.Error("expected non-zero exit code for cancelled context")
	}
}
