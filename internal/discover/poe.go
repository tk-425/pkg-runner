package discover

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type poeProvider struct{}

func (p *poeProvider) Name() string {
	return "poe"
}

func (p *poeProvider) Discover(dir string) ([]Script, error) {
	tomlPath := filepath.Join(dir, "pyproject.toml")
	if !fileExists(tomlPath) {
		return nil, nil
	}

	data, err := os.ReadFile(tomlPath)
	if err != nil {
		return nil, err
	}

	var config struct {
		Tool struct {
			Poe struct {
				Tasks map[string]string `toml:"tasks"`
			} `toml:"poe"`
		} `toml:"tool"`
	}

	if _, err := toml.Decode(string(data), &config); err != nil {
		return nil, err
	}

	tasks := config.Tool.Poe.Tasks
	if len(tasks) == 0 {
		return nil, nil
	}

	var scripts []Script
	for name := range tasks {
		scripts = append(scripts, Script{
			Name:    name,
			Command: "poe " + name,
			Source:  "poe",
		})
	}

	return scripts, nil
}
