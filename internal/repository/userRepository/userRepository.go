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

func (ur *UserRepository) GetUserByID(id string) (models.User,error){
	row := ur.Db.QueryRow("SELECT id, name, email, password, role FROM users WHERE id = ?", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (models.User,error){
	row := ur.Db.QueryRow("SELECT id, name, email, password, role FROM users WHERE email = ?", email)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (ur *UserRepository) GetUserCart(id string) (models.Cart,error){
	rows,err:=ur.Db.Query("SELECT cart FROM users WHERE id = ?",id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cart models.Cart
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		cart = append(cart, product)
	}
	return cart, nil
}

func (ur *UserRepository) AddToUserCart(userID string, product models.Product) error {
	_, err := ur.Db.Exec("UPDATE users SET cart = JSON_ARRAY_APPEND(cart, '$', JSON_OBJECT('id', ?, 'name', ?, 'price', ?, 'stock', ?)) WHERE id = ?",
		product.ID, product.Name, product.Price, product.Stock, userID)
	return err
}

func (ur *UserRepository) RemoveFromUserCart(userID string, product models.Product) error {
	_, err := ur.Db.Exec("UPDATE users SET cart = JSON_REMOVE(cart, JSON_UNQUOTE(JSON_SEARCH(cart, 'one', ?))) WHERE id = ?",
		product.ID, userID)
	return err
}