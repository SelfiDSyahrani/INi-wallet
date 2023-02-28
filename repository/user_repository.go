package repository

import (
<<<<<<< HEAD
	"dev_selfi/model"
	"dev_selfi/utils"
=======
	"INi-Wallet/model"
	"INi-Wallet/utils"
>>>>>>> development

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
<<<<<<< HEAD
=======
	Insert(user *model.User) error
	UpdateUserByID(id string) error
>>>>>>> development
	Delete(userWallet_ID string) error
	FindByEmail(email string) (*model.User, error)
	GetByID(userWallet_ID string) (model.User, error)
	GetAll() ([]model.User, error)
<<<<<<< HEAD
	Insert(user *model.User) (*model.User, error)
	UpdateById(user model.User) error
	
=======
	UpdateUserPass(id string, email string) error
>>>>>>> development
}

type userRepository struct {
	db *sqlx.DB
}

// Creates a new user
<<<<<<< HEAD
func (r *userRepository) Insert(user *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.UserWallet_ID)
=======
func (r *userRepository) Insert(newuser *model.User) error {
	_, err := r.db.NamedExec(utils.INSERT_USER, newuser)
>>>>>>> development
	if err != nil {
		return user, err
	}
	return user, nil
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

<<<<<<< HEAD
func (r *userRepository) UpdateById(user model.User) error {
	_,err := r.db.NamedExec(utils.UPDATE_USER_BY_ID,user)
=======
// UpdateUser updates an existing user
func (r *userRepository) UpdateUserByID(Id string) error {
	_, err := r.db.Exec(utils.UPDATE_USER_BY_ID, Id)
>>>>>>> development
	if err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	users := model.User{}
	err := r.db.Select(users, "SELECT * FROM users WHERE Email=$1", email)
	if err != nil {
		return nil, err
	}
	return &users, err
=======
// UpdateUser updates an existing user
func (r *userRepository) UpdateUserPass(Id string, email string) error {
	_, err := r.db.Exec(utils.UPDATE_USER_PASS, Id, email)
	if err != nil {
		return err
	}
	return nil
>>>>>>> development
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
