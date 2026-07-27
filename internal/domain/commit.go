package domain

type CommitType string

const (
	CommitTypeFeat     CommitType = "feat"
	CommitTypeFix      CommitType = "fix"
	CommitTypeDocs     CommitType = "docs"
	CommitTypeStyle    CommitType = "style"
	CommitTypeRefactor CommitType = "refactor"
	CommitTypeTest     CommitType = "test"
	CommitTypeChore    CommitType = "chore"
	CommitTypePerf     CommitType = "perf"
)

type Commit struct {
	Type        CommitType
	Scope       string
	Description string
	Body        string
}

func (c *Commit) FullMessage() string {
	if c.Scope != "" {
		return string(c.Type) + "(" + c.Scope + "): " + c.Description
	}
	return string(c.Type) + ": " + c.Description
}
