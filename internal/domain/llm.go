package domain

type LLMRequest struct {
	DiffText string
	Prompt   string
}

type LLMResponses struct {
	CommitType        string
	CommitScope       string
	CommitDescription string
	CommitBody        string
}

type LLMService interface {
	GenerateCommitMessage(req LLMRequest) (*LLMResponses, error)
}
