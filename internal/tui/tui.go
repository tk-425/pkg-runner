package tui

import (
	"context"
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tk-425/pkg-runner/internal/discover"
	"github.com/tk-425/pkg-runner/internal/runner"
)

type item struct {
	script discover.Script
}

func (i item) Title() string       { return i.script.Name }
func (i item) Description() string { return i.script.Command }
func (i item) FilterValue() string { return i.script.Name + " " + i.script.Command }

type exitStatusMsg struct{ code int }

type delegate struct{}

func (d delegate) Height() int                               { return 2 }
func (d delegate) Spacing() int                              { return 0 }
func (d delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d delegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	source := SourceStyle.Render("[" + i.script.Source + "]")
	name := NameStyle.Render(i.script.Name)
	cmd := CmdStyle.Render(i.script.Command)

	cursor := "  "
	if index == m.Index() {
		cursor = lipgloss.NewStyle().Foreground(lipgloss.Color("5")).Render("▸ ")
	}

	_, _ = fmt.Fprintf(w, "%s%s %s\n", cursor, source, name)
	_, _ = fmt.Fprintf(w, "   %s", cmd)
}

type model struct {
	list      list.Model
	runner    *runner.Runner
	quitting  bool
	executing bool
	status    *exitStatusMsg
	cancel    context.CancelFunc
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.executing && msg.String() == "ctrl+c" {
			m.cancel()
			m.executing = false
			return m, nil
		}
		if m.status != nil {
			m.status = nil
			return m, nil
		}
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if !m.executing {
				selected, ok := m.list.SelectedItem().(item)
				if ok {
					m.executing = true
					return m, runScript(selected.script, m.runner, &m.cancel)
				}
			}
		}

	case exitStatusMsg:
		m.executing = false
		m.status = &msg

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-2)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return ""
	}
	if m.status != nil {
		if m.status.code == 0 {
			return OkStyle.Render(fmt.Sprintf("\n  ✓ exit %d\n\n  press any key to continue", m.status.code))
		}
		return ErrStyle.Render(fmt.Sprintf("\n  ✗ exit %d\n\n  press any key to continue", m.status.code))
	}
	if len(m.list.Items()) == 0 {
		return HelpStyle.Render("No scripts found")
	}
	return m.list.View()
}

func runScript(script discover.Script, r *runner.Runner, cancel *context.CancelFunc) tea.Cmd {
	return func() tea.Msg {
		ctx, c := context.WithCancel(context.Background())
		*cancel = c
		defer c()
		code, err := r.Run(ctx, script)
		if err != nil {
			return exitStatusMsg{code: 1}
		}
		return exitStatusMsg{code: code}
	}
}

func NewProgram(scripts []discover.Script) *tea.Program {
	items := make([]list.Item, len(scripts))
	for i, s := range scripts {
		items[i] = item{script: s}
	}

	l := list.New(items, delegate{}, 40, 20)
	l.Title = "Scripts"
	l.SetShowHelp(true)
	l.SetFilteringEnabled(true)
	l.SetShowStatusBar(true)

	m := model{
		list:   l,
		runner: runner.New(),
	}

	return tea.NewProgram(m, tea.WithAltScreen())
}
