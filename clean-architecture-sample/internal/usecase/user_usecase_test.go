package usecase

import (
	"os"
	"testing"

	"clean-architecture-sample/internal/domain"
	"clean-architecture-sample/internal/infrastructure"
)

func setupTestRepo(t *testing.T) domain.UserRepository {
	path := "../../data/test_users.json"
	_ = os.Remove(path)
	return infrastructure.NewFileUserRepository(path)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	repo := setupTestRepo(t)
	uc := NewUserUseCase(repo)

	err := uc.CreateUser("1", "Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("CreateUser error: %v", err)
	}
}

func TestUserUseCase_GetUser(t *testing.T) {
	repo := setupTestRepo(t)
	uc := NewUserUseCase(repo)
	_ = uc.CreateUser("1", "Alice", "alice@example.com")

	user, err := uc.GetUser("1")
	if err != nil {
		t.Fatalf("GetUser error: %v", err)
	}
	if user.Name != "Alice" || user.Email != "alice@example.com" {
		t.Errorf("got %+v, want name=Alice, email=alice@example.com", user)
	}
}

func TestUserUseCase_UpdateUser(t *testing.T) {
	repo := setupTestRepo(t)
	uc := NewUserUseCase(repo)
	_ = uc.CreateUser("1", "Alice", "alice@example.com")

	err := uc.UpdateUser("1", "AliceUpdated", "alice2@example.com")
	if err != nil {
		t.Fatalf("UpdateUser error: %v", err)
	}
	user, err := uc.GetUser("1")
	if err != nil {
		t.Fatalf("GetUser error: %v", err)
	}
	if user.Name != "AliceUpdated" || user.Email != "alice2@example.com" {
		t.Errorf("got %+v, want name=AliceUpdated, email=alice2@example.com", user)
	}
}

func TestUserUseCase_DeleteUser(t *testing.T) {
	repo := setupTestRepo(t)
	uc := NewUserUseCase(repo)
	_ = uc.CreateUser("1", "Alice", "alice@example.com")

	err := uc.DeleteUser("1")
	if err != nil {
		t.Fatalf("DeleteUser error: %v", err)
	}
	_, err = uc.GetUser("1")
	if err == nil {
		t.Errorf("expected error for deleted user, got nil")
	}
}
