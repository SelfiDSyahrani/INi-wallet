package usecase

import (
	"INi-Wallet/model"
	"INi-Wallet/repository"
)

// User Use Case
type UserUseCase interface {
	RegisterUser(user model.User) error
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id string) error
}

// User Use Case implementation
type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) RegisterUser(user model.User) error {
	return u.userRepo.Create(user)
}

func (u *userUseCase) GetUserByID(id string) (model.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) GetAllUsers() ([]model.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) UpdateUser(user model.User) error {
	return u.userRepo.Update(user)
}

func (u *userUseCase) DeleteUser(id string) error {
	return u.userRepo.Delete(id)
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
