package repository

import (
	"INi-Wallet/model"
	"INi-Wallet/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Insert(user *model.User) error
	Update(user model.User) error
	Delete(userWallet_ID string) error
	GetByID(userWallet_ID string) (model.User, error)
	GetAll() ([]model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

// Creates a new user
func (r *userRepository) Insert(newuser model.User) error {
	_, err := r.db.NamedExec(utils.INSERT_USER, newuser)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing user
func (r *userRepository) Delete(user_ID string) error {
	_, err := r.db.Exec(utils.DELETE_PAYMENT_METHOD, user_ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a user by ID
//GetByID
func (r *userRepository) GetByID(user_ID string) (model.user, error) {
	var user model.User
	err := r.db.QueryRow(utils.SELECT_USER_ID, user_ID).Scan(
		&user.UserWallet_ID,
		&user.name,
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
func (r *userRepository) Update(users model.User) error {
	_, err := r.db.Exec(utils.UPDATE_USER, users)
	return r.db.Save(users).Error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
