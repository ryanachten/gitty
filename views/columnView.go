package views

import (
	"fmt"

	services "github.com/ryanachten/gitty/services"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	modelStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("99")).
			Padding(1)
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
)

type model struct {
	results services.GittyResults
}

func ColumnView(results services.GittyResults) model {
	return model{
		results,
	}
}

func (s model) Init() tea.Cmd { return nil }

func (m model) View() string {
	resultCount := len(m.results)

	cells := make([]string, resultCount)
	index := 0

	for repo, output := range m.results {

		var content string
		if output.Err != nil {
			content = output.Err.Error()
		} else {
			content = output.Result
		}

		header := fmt.Sprintln(headerStyle.Render(repo.Label))
		cell := fmt.Sprintf("\n%v", header) + modelStyle.Render(content)
		cells[index] = cell

		index++
	}

	contents := lipgloss.JoinVertical(lipgloss.Top, cells...)
	contents += helpStyle.Render(fmt.Sprintln("\nq: exit"))

	return contents
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}
