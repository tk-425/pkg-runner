package discover

import (
	"errors"
	"os"
)

type Script struct {
	Name    string
	Command string
	Source  string
}

type Provider interface {
	Name() string
	Discover(dir string) ([]Script, error)
}

var providers = []Provider{
	&npmProvider{},
	&makefileProvider{},
	&poeProvider{},
	&cargoProvider{},
	&goProvider{},
}

func Discover(dir string) ([]Script, error) {
	var allScripts []Script
	var errs []error

	for _, p := range providers {
		scripts, err := p.Discover(dir)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		allScripts = append(allScripts, scripts...)
	}

	return allScripts, errors.Join(errs...)
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}
