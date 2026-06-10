package discover

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type npmProvider struct{}

func (p *npmProvider) Name() string {
	return "npm"
}

func (p *npmProvider) Discover(dir string) ([]Script, error) {
	pkgPath := filepath.Join(dir, "package.json")
	if !fileExists(pkgPath) {
		return nil, nil
	}

	data, err := os.ReadFile(pkgPath)
	if err != nil {
		return nil, err
	}

	var pkg struct {
		Scripts map[string]string `json:"scripts"`
	}
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	if len(pkg.Scripts) == 0 {
		return nil, nil
	}

	runner := detectRunner(dir)

	var scripts []Script
	for name := range pkg.Scripts {
		scripts = append(scripts, Script{
			Name:    name,
			Command: runner + " run " + name,
			Source:  "npm",
		})
	}

	return scripts, nil
}

func detectRunner(dir string) string {
	switch {
	case fileExists(filepath.Join(dir, "pnpm-lock.yaml")):
		return "pnpm"
	case fileExists(filepath.Join(dir, "yarn.lock")):
		return "yarn"
	case fileExists(filepath.Join(dir, "bun.lockb")):
		return "bun"
	default:
		return "npm"
	}
}
