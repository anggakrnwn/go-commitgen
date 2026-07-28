package repository

import (
	"fmt"

	"github.com/anggakrnwn/go-commitgen/internal/domain"
	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/git"
)

type GitRepository struct {
	client *git.Client
}

func NewGitRepository(client *git.Client) domain.GitRepository {
	return &GitRepository{client: client}

}

func (r *GitRepository) GetStagedDiff() (string, error) {

	diff, err := r.client.DiffCached()
	if err != nil {
		return "", err
	}

	if diff == "" {
		return "", fmt.Errorf("no staged changes found. please run 'git add' first")
	}

	return diff, nil
}
