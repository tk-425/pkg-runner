# pkg-runner

Multi-language script discovery and execution — find and run scripts from `package.json`, `Makefile`, `pyproject.toml`, `Cargo.toml`, and Go with a colorful terminal UI.

## Features

- **Auto-discovery** — detects all runnable scripts in the current directory from supported manifests
- **Interactive TUI** — arrow-key selector with fuzzy search, powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- **Source-tagged** — each script shows its origin (`[npm]`, `[make]`, `[poe]`, `[cargo]`, `[go]`)
- **Cross-platform** — Linux, macOS, Windows (amd64, arm64)

## Supported Sources

| Source | Manifest |
|--------|----------|
| `[npm]` | `package.json` scripts |
| `[make]` | `Makefile` targets |
| `[poe]` | `pyproject.toml` [tool.poe.tasks] |
| `[cargo]` | `Cargo.toml` |
| `[go]` | Go tool invocations |

## Install

```bash
go install github.com/tk-425/pkg-runner/cmd/pkg-runner@latest
```

Or download a prebuilt binary from [releases](https://github.com/tk-425/pkg-runner/releases).

## Usage

```bash
# Interactive mode — arrow keys to select, enter to run, esc to quit
pkg-runner

# Print version
pkg-runner --version
pkg-runner -v
```

## Build

```bash
go build -ldflags "-X main.version=$(git describe --tags --abbrev=0)" ./cmd/pkg-runner
```
