package discover

import (
	"path/filepath"
)

type goProvider struct{}

func (p *goProvider) Name() string {
	return "go"
}

func (p *goProvider) Discover(dir string) ([]Script, error) {
	goModPath := filepath.Join(dir, "go.mod")
	if !fileExists(goModPath) {
		return nil, nil
	}

	makefileProvider := &makefileProvider{}
	scripts, err := makefileProvider.Discover(dir)
	if err != nil {
		return nil, err
	}

	for i := range scripts {
		scripts[i].Source = "go"
		scripts[i].Command = "make " + scripts[i].Name
	}

	return scripts, nil
}
