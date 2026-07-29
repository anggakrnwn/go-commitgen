package main

import (
	"fmt"
	"os"

	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/config"
	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/git"
	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/llm"
	"github.com/anggakrnwn/go-commitgen/internal/repository"
	"github.com/anggakrnwn/go-commitgen/internal/usecase"
)

func main() {
	cfg := config.Connection()
	gitClint := git.NewClient()

	geminiClient, err := llm.NewGeminiClient(cfg.GeminiAPIKey, cfg.DefaultModel)
	if err != nil {
		fmt.Printf("error initializing model: %v\n", err)
		os.Exit(1)
	}

	gitRepo := repository.NewGitRepository(gitClint)

	commitUsecase := usecase.NewCommitUsecase(gitRepo, geminiClient)
	commitMsg, err := commitUsecase.GenerateCommit()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("generate commit message:\n%s\n", commitMsg)
}
