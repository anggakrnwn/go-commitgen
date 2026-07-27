package git

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/anggakrnwn/go-commitgen/internal/domain"
)

type Service struct{}

func NewGitService() domain.GitRepository {
	return &Service{}
}

func (s *Service) GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute git diff: %v, stderr: %s", err, stderr.String())
	}

	diffOutput := out.String()
	if diffOutput == "" {
		return "", fmt.Errorf("no staged changes found. please run 'git add' first")
	}

	return diffOutput, nil

}
