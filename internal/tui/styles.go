package tui

import "github.com/charmbracelet/lipgloss"

var (
	SourceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	NameStyle   = lipgloss.NewStyle().Bold(true)
	CmdStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Faint(true)
	OkStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("2")).Bold(true)
	ErrStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Bold(true)
	HelpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)
