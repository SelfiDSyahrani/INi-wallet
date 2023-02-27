package usecase

import (
	"INi-Wallet/dto"
	"INi-Wallet/model"
	"INi-Wallet/repository"
	"net/mail"

	"gorm.io/gorm/utils"
)

// User Use Case
type UserUseCase interface {
	RegisterUser(input *dto.UserRequestBody) (model.User, error)
	GetUserByID(id string) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id string) error
}

// User Use Case implementation
type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) RegisterUser(input *dto.UserRequestBody) (model.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return model.User{}, err
	}
	user, err := u.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.ID != 0 {
		return user, utils.custom_error.UserAlreadyExistsError{}
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}


func (u *userUseCase) GetUserByID(id string) (*model.User, error) {
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
