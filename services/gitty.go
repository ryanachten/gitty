package services

import (
	"fmt"
	"os/exec"

	"gitty/models"
)

type Gitty struct {
	repositories []models.Repository
}

// Initialises a new Gitty service
func CreateGitty(repositories []models.Repository) Gitty {
	gitty := Gitty{
		repositories,
	}

	return gitty
}

// Runs arbitrary git commands in each folder
func (g *Gitty) Run(args []string) {
	for _, repo := range g.repositories {

		cmd := exec.Command("git", args...)
		cmd.Dir = repo.WorkingDirectory

		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error running command: %v", err)
		}

		fmt.Printf("Output: \n %v", string(output))
	}
}
