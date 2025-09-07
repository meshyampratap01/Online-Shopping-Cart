package userRepository

import (
	"database/sql"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserManager {
	return &UserRepository{Db: db}
}

func (ur *UserRepository) SaveUser(user models.User) error {
	_, err := ur.Db.Exec("INSERT INTO users (id, name, email, password, role) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Email, user.Password, user.Role)
	return err
}

func (ur *UserRepository) GetUserByID(id string) (models.User, error) {
	row := ur.Db.QueryRow("SELECT id, name, email, password, role FROM users WHERE id = ?", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User, error) {
	row := ur.Db.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = ?", email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
