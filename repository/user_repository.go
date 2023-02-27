package repository

import (
	"INi-Wallet/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user model.User) error
	Update(user model.User) error
	Delete(userWallet_ID string) error
	GetByID(userWallet_ID string) (*model.User, error)
	GetAll() ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// Creates a new user
func (r *userRepository) Create(user model.User) error {
	// Use GORM to save the user
	return r.db.Create(user).Error
}

// DeleteUser deletes an existing user
func (r *userRepository) Delete(id string) error {
	// Use GORM to delete the user
	user := model.User{}
	user.UserWallet_ID = id
	return r.db.Where("id = ?", id).Delete(&user).Error
}

// GetByID retrieves a user by ID
func (r *userRepository) GetByID(id string) (*model.User, error) {
	// Use GORM to retrieve the user
	user := model.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

// GetAll retrieves all users
func (r *userRepository) GetAll() ([]model.User, error) {
	// Use GORM to retrieve all users
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

// UpdateUser updates an existing user
func (r *userRepository) Update(user model.User) error {
	// Use GORM to update the user
	return r.db.Save(user).Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

