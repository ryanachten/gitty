package views

import (
	"fmt"
	"os"

	services "gitty/services"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

var (
	modelStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder())
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
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
	fd := int(os.Stdout.Fd())
	width, _, _ := term.GetSize(fd)

	resultCount := len(m.results)

	cellWidth := width/resultCount - 4*resultCount

	cells := make([]string, resultCount)
	index := 0

	for _, output := range m.results {

		var content string
		if output.Err != nil {
			content = output.Err.Error()
		} else {
			content = output.Result
		}

		wrappedContent := wordwrap.String(content, cellWidth)
		cell := modelStyle.Width(cellWidth).Render(wrappedContent)
		cells[index] = cell

		index++
	}

	contents := lipgloss.JoinVertical(lipgloss.Top, cells...)
	contents += helpStyle.Render(fmt.Sprintln("\nctrl + c: exit"))

	return contents
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		switch msg.(tea.KeyMsg).String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}
