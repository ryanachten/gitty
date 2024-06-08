package main

import (
	"log"

	"github.com/ryanachten/gitty/models"
	"github.com/ryanachten/gitty/services"
	"github.com/ryanachten/gitty/views"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	arguments := models.GetArguments()

	config, err := models.ParseConfigurationFile(arguments.ConfigurationPath)
	if err != nil {
		log.Fatalln(err)
	}

	gitty := services.CreateGitty(config.Repositories)
	results := gitty.Run(arguments.Command)

	p := tea.NewProgram(
		views.ColumnView(results),
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
