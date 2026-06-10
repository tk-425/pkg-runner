package discover

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type makefileProvider struct{}

func (p *makefileProvider) Name() string {
	return "make"
}

func (p *makefileProvider) Discover(dir string) ([]Script, error) {
	makefilePath := filepath.Join(dir, "Makefile")
	f, err := os.Open(makefilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer func() { _ = f.Close() }()

	var scripts []Script
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		target := parseTarget(line)
		if target == "" {
			continue
		}
		if strings.HasPrefix(target, ".") {
			continue
		}
		scripts = append(scripts, Script{
			Name:    target,
			Command: "make " + target,
			Source:  "make",
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return scripts, nil
}

func parseTarget(line string) string {
	if line == "" {
		return ""
	}
	if line[0] == '\t' || line[0] == ' ' || line[0] == '#' {
		return ""
	}

	colon := strings.Index(line, ":")
	if colon < 0 {
		return ""
	}

	target := strings.TrimSpace(line[:colon])
	if target == "" {
		return ""
	}

	return strings.Fields(target)[0]
}
