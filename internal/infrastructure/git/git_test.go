package git_test

import (
	"testing"

	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/git"
)

func TestService_GetStagedDiff(t *testing.T) {
	gitService := git.NewGitService()
	tests := []struct {
		name string // description of this test case
		// want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test GetStagedDiff execution",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			got, gotErr := gitService.GetStagedDiff()
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetStagedDiff() error: %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("GetStagedDiff() got empty string, expected diff output")
			}
		})
	}
}
