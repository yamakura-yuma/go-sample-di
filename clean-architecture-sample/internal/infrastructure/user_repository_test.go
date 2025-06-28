package infrastructure

import (
	"clean-architecture-sample/internal/domain"
	"os"
	"testing"
)

func TestFileUserRepository_SaveAndFindByID(t *testing.T) {
	repo := NewFileUserRepository("../../data/test_users_repo.json")
	_ = os.Remove("../../data/test_users_repo.json")
	user := &domain.User{ID: "1", Name: "RepoTest", Email: "repo@example.com"}
	if err := repo.Save(user); err != nil {
		t.Fatalf("Save error: %v", err)
	}
	got, err := repo.FindByID("1")
	if err != nil {
		t.Fatalf("FindByID error: %v", err)
	}
	if got.Name != user.Name || got.Email != user.Email {
		t.Errorf("got %+v, want %+v", got, user)
	}
}
