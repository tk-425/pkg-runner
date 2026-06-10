package discover

import (
	"testing"
)

func TestMakefileProvider_SpecialTargetsOnly(t *testing.T) {
	provider := &makefileProvider{}
	scripts, err := provider.Discover("testdata/special-only")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(scripts) != 0 {
		t.Fatalf("expected 0 scripts, got %d", len(scripts))
	}
}

func TestMakefileProvider_Discover(t *testing.T) {
	provider := &makefileProvider{}
	scripts, err := provider.Discover("testdata")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(scripts) != 3 {
		t.Fatalf("expected 3 scripts, got %d", len(scripts))
	}

	expected := map[string]string{
		"build": "make build",
		"test":  "make test",
		"clean": "make clean",
	}

	for _, s := range scripts {
		wantCmd, ok := expected[s.Name]
		if !ok {
			t.Errorf("unexpected script: %s", s.Name)
			continue
		}
		if s.Command != wantCmd {
			t.Errorf("script %q: expected command %q, got %q", s.Name, wantCmd, s.Command)
		}
		if s.Source != "make" {
			t.Errorf("script %q: expected source %q, got %q", s.Name, "make", s.Source)
		}
	}
}

func TestMakefileProvider_MissingMakefile(t *testing.T) {
	provider := &makefileProvider{}
	scripts, err := provider.Discover("testdata/missing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if scripts != nil {
		t.Fatalf("expected nil scripts, got %v", scripts)
	}
}

func TestMakefileProvider_EmptyMakefile(t *testing.T) {
	provider := &makefileProvider{}
	scripts, err := provider.Discover("testdata/empty-makefile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(scripts) != 0 {
		t.Fatalf("expected 0 scripts, got %d", len(scripts))
	}
}

// Poe provider tests

func TestPoeProvider_Discover(t *testing.T) {
	provider := &poeProvider{}
	scripts, err := provider.Discover("testdata")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(scripts) != 3 {
		t.Fatalf("expected 3 scripts, got %d", len(scripts))
	}

	expected := map[string]string{
		"build": "poe build",
		"test":  "poe test",
		"lint":  "poe lint",
	}

	for _, s := range scripts {
		wantCmd, ok := expected[s.Name]
		if !ok {
			t.Errorf("unexpected script: %s", s.Name)
			continue
		}
		if s.Command != wantCmd {
			t.Errorf("script %q: expected command %q, got %q", s.Name, wantCmd, s.Command)
		}
		if s.Source != "poe" {
			t.Errorf("script %q: expected source %q, got %q", s.Name, "poe", s.Source)
		}
	}
}

func TestPoeProvider_NoPoeSection(t *testing.T) {
	provider := &poeProvider{}
	scripts, err := provider.Discover("testdata/no-poe")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(scripts) != 0 {
		t.Fatalf("expected 0 scripts, got %d", len(scripts))
	}
}

func TestPoeProvider_MissingFile(t *testing.T) {
	provider := &poeProvider{}
	scripts, err := provider.Discover("testdata/missing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if scripts != nil {
		t.Fatalf("expected nil scripts, got %v", scripts)
	}
}

func TestPoeProvider_MalformedTOML(t *testing.T) {
	provider := &poeProvider{}
	_, err := provider.Discover("testdata/malformed-toml")
	if err == nil {
		t.Fatal("expected error for malformed TOML, got nil")
	}
}

// NPM provider tests

func TestNpmProvider_Discover(t *testing.T) {
	provider := &npmProvider{}
	scripts, err := provider.Discover("testdata")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(scripts) != 3 {
		t.Fatalf("expected 3 scripts, got %d", len(scripts))
	}

	expected := map[string]string{
		"build": "npm run build",
		"test":  "npm run test",
		"lint":  "npm run lint",
	}

	for _, s := range scripts {
		wantCmd, ok := expected[s.Name]
		if !ok {
			t.Errorf("unexpected script: %s", s.Name)
			continue
		}
		if s.Command != wantCmd {
			t.Errorf("script %q: expected command %q, got %q", s.Name, wantCmd, s.Command)
		}
		if s.Source != "npm" {
			t.Errorf("script %q: expected source %q, got %q", s.Name, "npm", s.Source)
		}
	}
}

func TestNpmProvider_MissingFile(t *testing.T) {
	provider := &npmProvider{}
	scripts, err := provider.Discover("testdata/missing")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if scripts != nil {
		t.Fatalf("expected nil scripts, got %v", scripts)
	}
}

func TestNpmProvider_MalformedJSON(t *testing.T) {
	provider := &npmProvider{}
	_, err := provider.Discover("testdata/malformed-json")
	if err == nil {
		t.Fatal("expected error for malformed JSON, got nil")
	}
}

// Cargo and Go providers

func TestCargoProvider_WithMakefile(t *testing.T) {
	provider := &cargoProvider{}
	scripts, err := provider.Discover("testdata/cargo-with-makefile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(scripts) != 2 {
		t.Fatalf("expected 2 scripts, got %d", len(scripts))
	}

	for _, s := range scripts {
		if s.Source != "cargo" {
			t.Errorf("script %q: expected source cargo, got %s", s.Name, s.Source)
		}
	}
}

func TestCargoProvider_WithoutMakefile(t *testing.T) {
	provider := &cargoProvider{}
	scripts, err := provider.Discover("testdata/cargo-only")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(scripts) != 0 {
		t.Fatalf("expected 0 scripts, got %d", len(scripts))
	}
}

func TestGoProvider_WithMakefile(t *testing.T) {
	provider := &goProvider{}
	scripts, err := provider.Discover("testdata/go-with-makefile")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(scripts) != 2 {
		t.Fatalf("expected 2 scripts, got %d", len(scripts))
	}

	for _, s := range scripts {
		if s.Source != "go" {
			t.Errorf("script %q: expected source go, got %s", s.Name, s.Source)
		}
	}
}

func TestGoProvider_WithoutMakefile(t *testing.T) {
	provider := &goProvider{}
	scripts, err := provider.Discover("testdata/go-only")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(scripts) != 0 {
		t.Fatalf("expected 0 scripts, got %d", len(scripts))
	}
}
