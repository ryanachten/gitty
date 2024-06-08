package services

import (
	"os/exec"

	"gitty/models"
)

// Performs git operations in a list of local repositories.
// Relies on native git installations due to issues with git Go module CRLF support on Windows
type Gitty struct {
	repositories []models.Repository
}

type GittyOutput struct {
	Result string
	Err    error
}

type GittyResults map[models.Repository]GittyOutput

// Initialises a new Gitty service
func CreateGitty(repositories []models.Repository) Gitty {
	gitty := Gitty{
		repositories,
	}

	return gitty
}

// Runs arbitrary git commands in each folder
func (g *Gitty) Run(args []string) GittyResults {
	result := make(GittyResults)

	for _, repo := range g.repositories {
		cmd := exec.Command("git", args...)
		cmd.Dir = repo.WorkingDirectory

		output, err := cmd.Output()

		result[repo] = GittyOutput{
			Err:    err,
			Result: string(output),
		}
	}

	return result
}
