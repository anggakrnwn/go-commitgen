package usecase

import (
	"fmt"
	"strings"

	"github.com/anggakrnwn/go-commitgen/internal/domain"
)

type CommitUsecase struct {
	gitRepo domain.GitRepository
	llmSvc  domain.LLMService
}

func NewCommitUsecase(gitRepo domain.GitRepository, llmSvc domain.LLMService) *CommitUsecase {
	return &CommitUsecase{
		gitRepo: gitRepo,
		llmSvc:  llmSvc,
	}

}

func (u *CommitUsecase) GenerateCommit() (string, error) {

	diff, err := u.gitRepo.GetStagedDiff()
	if err != nil {
		return "", err
	}

	textLLM := "Analyze the following git diff and generate a conventional commit message. " +
		"Provide the commit type, optional scope, a concise description, and an optional detailed body."

	req := domain.LLMRequest{
		DiffText: diff,
		Prompt:   textLLM,
	}

	llmResp, err := u.llmSvc.GenerateCommitMessage(req)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message from AI: %w", err)
	}

	var sb strings.Builder
	sb.WriteString(llmResp.CommitType)

	if llmResp.CommitScope != "" {
		sb.WriteString(fmt.Sprintf("(%s)", llmResp.CommitScope))
	}

	sb.WriteString(fmt.Sprintf(": %s", llmResp.CommitDescription))

	if llmResp.CommitBody != "" {
		sb.WriteString(fmt.Sprintf("\n\n%s", llmResp.CommitBody))
	}

	return sb.String(), nil

}
