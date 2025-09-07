package cartRepository

import (
	"database/sql"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type CartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartManager {
	return &CartRepository{db: db}
}

func (cr *CartRepository) CreateCart(cartID, userID string) error {
	_, err := cr.db.Exec("INSERT INTO cart (id, user_id) VALUES (?,?)", cartID, userID)
	return err
}

func (cr *CartRepository) AddToCart(userID string, product models.Product) error {
	cartID, err := cr.GetCartIDByUserID(userID)
	if err != nil {
		return err
	}
	row := cr.db.QueryRow("SELECT product_id FROM cart_items WHERE cart_id=? AND product_id=?", cartID, product.ID)
	var product_id string
	if err := row.Scan(&product_id); err != nil {
		if err == sql.ErrNoRows {
			_, err = cr.db.Exec("INSERT INTO cart_items (cart_id, product_id, quantity) VALUES (?, ?, ?)", cartID, product.ID, 1)
			return err
		}
		return err
	}
	_, err = cr.db.Exec("UPDATE cart_items SET quantity = quantity + 1 WHERE product_id = ?", product_id)
	return err
}

func (cr *CartRepository) RemoveFromCart(userID string, prodID string) error {
	cartID, err := cr.GetCartIDByUserID(userID)
	if err != nil {
		return err
	}
	_, err = cr.db.Exec("DELETE FROM cart_items WHERE cart_id = ? AND product_id = ?", cartID, prodID)
	return err
}

func (cr *CartRepository) GetCartIDByUserID(userID string) (string, error) {
	row := cr.db.QueryRow("SELECT id FROM cart WHERE user_id = ?", userID)
	var cartID string
	err := row.Scan(&cartID)
	if err != nil {
		return "", err
	}
	return cartID, nil
}

func (cr *CartRepository) EmptyCart(userID string) error {
	cartID, err := cr.GetCartIDByUserID(userID)
	if err != nil {
		return err
	}
	_, err = cr.db.Exec("DELETE FROM cart_items WHERE cart_id = ?", cartID)
	return err
}

func (cr *CartRepository) GetCartItems(cartID string) ([]dto.CartItemsDTO, error) {
	rows, err := cr.db.Query(`
		SELECT p.id, p.name, p.price, ci.quantity
		FROM cart_items ci
		JOIN products p ON ci.product_id = p.id
		WHERE ci.cart_id = ?`, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cartItems []dto.CartItemsDTO
	for rows.Next() {
		var item dto.CartItemsDTO
		err := rows.Scan(&item.ProductID, &item.ProductName, &item.Price, &item.Quantity)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, item)
	}
	return cartItems, nil
}
