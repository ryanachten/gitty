package main

import (
	"log"

	"gitty/models"
	"gitty/services"
)

func main() {
	arguments := models.GetArguments()

	config, err := models.ParseConfigurationFile(arguments.ConfigurationPath)
	if err != nil {
		log.Fatalln(err)
	}

	gitty := services.CreateGitty(config.Repositories)
	gitty.Run(arguments.Command)

	// p := tea.NewProgram(
	// 	views.NewSimplePage("This app is under construction"),
	// )

	// if _, err := p.Run(); err != nil {
	// 	panic(err)
	// }
}
