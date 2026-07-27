package domain

type GitDiff struct {
	DiffContent string
	FileChange  []string
}

type GitRepository interface {
	GetStagedDiff() (string, error)
}
