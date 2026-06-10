package discover

import (
	"path/filepath"
)

type cargoProvider struct{}

func (p *cargoProvider) Name() string {
	return "cargo"
}

func (p *cargoProvider) Discover(dir string) ([]Script, error) {
	cargoPath := filepath.Join(dir, "Cargo.toml")
	if !fileExists(cargoPath) {
		return nil, nil
	}

	makefileProvider := &makefileProvider{}
	scripts, err := makefileProvider.Discover(dir)
	if err != nil {
		return nil, err
	}

	for i := range scripts {
		scripts[i].Source = "cargo"
		scripts[i].Command = "make " + scripts[i].Name
	}

	return scripts, nil
}
