package repository

type UserRepository interface {
	Delete(userWallet_ID string) error
	FindByEmail(email string) (*model.User, error)
	GetByID(userWallet_ID string) (model.User, error)
	GetAll() ([]model.User, error)
	Insert(user *model.User) (*model.User, error)
	UpdateById(user model.User) error
}

type userRepository struct {
	db *sqlx.DB
}

// Creates a new user
func (r *userRepository) Insert(user *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.UserWallet_ID)
	if err != nil {
		return user, err
	}
	return user, nil
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

func (r *userRepository) UpdateById(user model.User) error {
	_, err := r.db.NamedExec(utils.UPDATE_USER_BY_ID, user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	users := model.User{}
	err := r.db.Select(users, "SELECT * FROM user WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &users, err
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
