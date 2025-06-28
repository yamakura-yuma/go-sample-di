package usecase

import "clean-architecture-sample/internal/domain"

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(id, name, email string) error {
	user := &domain.User{ID: id, Name: name, Email: email}
	return u.repo.Save(user)
}

func (u *UserUseCase) GetUser(id string) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *UserUseCase) UpdateUser(id, name, email string) error {
	user := &domain.User{ID: id, Name: name, Email: email}
	return u.repo.Update(user)
}

func (u *UserUseCase) DeleteUser(id string) error {
	return u.repo.Delete(id)
}
