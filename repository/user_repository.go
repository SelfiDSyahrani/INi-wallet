package repository

import (
	"INi-Wallet/model"
	"INi-Wallet/utils"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Insert(user *model.User) error
	UpdateUserByID(id string) error
	Delete(userWallet_ID string) error
	GetByID(userWallet_ID string) (model.User, error)
	GetAll() ([]model.User, error)
	UpdateUserPass(id string, email string) error
}

type userRepository struct {
	db *sqlx.DB
}

// Creates a new user
func (r *userRepository) Insert(newuser *model.User) error {
	_, err := r.db.NamedExec(utils.INSERT_USER, newuser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing user
func (r *userRepository) Delete(user_ID string) error {
	_, err := r.db.Exec(utils.DELETE_USER, user_ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a user by ID
// GetByID
func (r *userRepository) GetByID(user_ID string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(utils.SELECT_USER_ID, user_ID).Scan(
		&user.UserWallet_ID,
		&user.Name,
	)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetAll retrieves all users
func (r *userRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Select(users, utils.SELECT_USER_LIST)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates an existing user
func (r *userRepository) UpdateUserByID(Id string) error {
	_, err := r.db.Exec(utils.UPDATE_USER_BY_ID, Id)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates an existing user
func (r *userRepository) UpdateUserPass(Id string, email string) error {
	_, err := r.db.Exec(utils.UPDATE_USER_PASS, Id, email)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
