package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/raikou/internal/ssh"
	"github.com/raikou/internal/utils"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Running command: %s\n", utils.CreateSSHCommand(
					m.table.SelectedRow()[0],
					m.table.SelectedRow()[1],
					m.table.SelectedRow()[2],
					123,
					m.table.SelectedRow()[4],
				)),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func RenderTable(hosts *[]ssh.SSHosts) {
	columns := []table.Column{
		{Title: "Host", Width: 4},
		{Title: "Hostname", Width: 10},
		{Title: "User", Width: 10},
		{Title: "Port", Width: 10},
		{Title: "IdentityFile", Width: 10},
	}

	var rows []table.Row

	for _, host := range *hosts {
		rows = append(rows, table.Row{
			host.Host, host.HostName, host.User, fmt.Sprintf("%d", host.Port), host.IdentityFile,
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
