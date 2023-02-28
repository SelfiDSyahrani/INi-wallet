package usecase

import (
	"dev_selfi/model"
	"dev_selfi/repository"
	"dev_selfi/utils"
	"net/mail"

	"dev_selfi/dto"

	"golang.org/x/crypto/bcrypt"
)

// User Use Case
type UserUseCase interface {
	RegisterUser(input *dto.UserRequestBody) (*model.User, error)
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id string) error
}

// User Use Case implementation
type userUseCase struct {
	userRepo repository.UserRepository
}

type USConfig struct {
	UserRepository repository.UserRepository
}

func (u *userUseCase) RegisterUser(input *dto.UserRequestBody) (*model.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &model.User{}, err
	}
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return user, &utils.NotValidEmailError{}
	}
	if user.Email == input.Email {
		return user, &utils.UserAlreadyExistsError{}
	}
	user.UserWallet_ID = utils.GenerateId()
	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := u.userRepo.Insert(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (u *userUseCase) GetUserByID(id string) (model.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) GetAllUsers() ([]model.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) UpdateUser(user model.User) error {
	return u.userRepo.UpdateById(user)
}

func (u *userUseCase) DeleteUser(id string) error {
	return u.userRepo.Delete(id)
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
