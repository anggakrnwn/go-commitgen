package repository_test

import (
	"testing"

	"github.com/anggakrnwn/go-commitgen/internal/infrastructure/git"
	"github.com/anggakrnwn/go-commitgen/internal/repository"
)

func TestGitRepository_GetStagedDiff(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "skenario gagal atau kosong saat tidak ada staged files",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			infraClient := git.NewClient()

			repo := repository.NewGitRepository(infraClient)

			got, gotErr := repo.GetStagedDiff()
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("GetStagedDiff() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("GetStagedDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
