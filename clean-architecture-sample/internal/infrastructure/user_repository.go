package infrastructure

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"clean-architecture-sample/internal/domain"
)

// ファイル保存用の実装
type FileUserRepository struct {
	filePath string
	mu       sync.Mutex
}

func NewFileUserRepository(filePath string) *FileUserRepository {
	return &FileUserRepository{filePath: filePath}
}

func (r *FileUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	users, _ := r.loadAll()
	users = append(users, *user)
	return r.saveAll(users)
}

func (r *FileUserRepository) FindByID(id string) (*domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	users, _ := r.loadAll()
	for _, u := range users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, os.ErrNotExist
}

func (r *FileUserRepository) loadAll() ([]domain.User, error) {
	file, err := os.OpenFile(r.filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil || len(bytes) == 0 {
		return []domain.User{}, nil
	}
	var users []domain.User
	if err := json.Unmarshal(bytes, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *FileUserRepository) saveAll(users []domain.User) error {
	bytes, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, bytes, 0666)
}

func (r *FileUserRepository) Update(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	users, _ := r.loadAll()
	updated := false
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			updated = true
			break
		}
	}
	if !updated {
		return os.ErrNotExist
	}
	return r.saveAll(users)
}

func (r *FileUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	users, _ := r.loadAll()
	newUsers := make([]domain.User, 0, len(users))
	deleted := false
	for _, u := range users {
		if u.ID == id {
			deleted = true
			continue
		}
		newUsers = append(newUsers, u)
	}
	if !deleted {
		return os.ErrNotExist
	}
	return r.saveAll(newUsers)
}

// インターフェース用
var _ domain.UserRepository = (*FileUserRepository)(nil)
