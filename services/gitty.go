package services

import (
	"fmt"
	"os/exec"

	"gitty/models"

	"github.com/go-git/go-git/v5"
)

type Gitty struct {
	projects     []models.Project
	repositories []*git.Repository
}

// Initialises a new Gitty service
func CreateGitty(projects []models.Project) Gitty {
	gitty := Gitty{
		projects: projects,
	}

	for _, project := range gitty.projects {
		repository, err := git.PlainOpen(project.Location)
		if err != nil {
			fmt.Printf("Error opening repository: %v\n", err)
		} else {
			gitty.repositories = append(gitty.repositories, repository)
		}
	}

	return gitty
}

func (g *Gitty) Status() {
	for _, repo := range g.repositories {
		worktree, err := repo.Worktree()
		if err != nil {
			fmt.Printf("Error getting worktree: %v\n", err)
			return
		}
		status, err := worktree.Status()
		if err != nil {
			fmt.Printf("Error getting status: %v\n", err)
			return
		}
		worktree.Status()

		fmt.Println(status)

		cmd := exec.Command("git", "status", "--porcelain")
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("could not get status from git command-line: %v", err)
		}

		// Print the status from the Git command-line tool
		fmt.Println("git command-line repository status:")
		fmt.Println(string(output))
	}
}
