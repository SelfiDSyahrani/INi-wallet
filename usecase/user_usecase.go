package usecase

import (
	"INi-Wallet/dto"
	"INi-Wallet/model"
	"INi-Wallet/repository"
	"INi-Wallet/utils"

	// "net/mail"

	"golang.org/x/crypto/bcrypt"
)

// User Use Case
type UserUseCase interface {
	RegisterUser(input *model.User) error
	GetByEmail(email string) (model.User, error)
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id string) error
	Login(input *dto.LoginRequestBody) (model.User, error)
}

// User Use Case implementation
type userUseCase struct {
	userRepo repository.UserRepository
}

type USConfig struct {
	UserRepository repository.UserRepository
}

func (u *userUseCase) RegisterUser(input *model.User) error {
	if _, err := u.userRepo.FindByEmail(input.Email); err == nil {
		return &utils.UserAlreadyExistsError{}
	}
	input.ID = utils.GenerateId()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	input.Password = string(passwordHash)
	input.Balance = 0
	return u.userRepo.Insert(input)
}

func (u *userUseCase) GetByEmail(email string) (model.User, error) {
	return u.userRepo.FindByEmail(email)
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

func (s *userUseCase) Login(input *dto.LoginRequestBody) (model.User, error) {
	// _, err := mail.ParseAddress(input.Email)
	// if err != nil {
	// 	return model.User{}, &utils.NotValidEmailError{}
	// }
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, &utils.IncorrectCredentialsError{}
	}
	return user, nil
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
