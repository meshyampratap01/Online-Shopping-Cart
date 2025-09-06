package productRepository

import (
	"database/sql"

	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductManager {
	return &ProductRepository{Db: db}
}

func (pr *ProductRepository) AddProduct(product models.Product) error {
	_, err := pr.Db.Exec("INSERT INTO products (id, name, price, stock) VALUES (?, ?, ?, ?)",
		product.ID, product.Name, product.Price, product.Stock)
	return err
}

func (pr *ProductRepository) RemoveProduct(id string) error {
	_, err := pr.Db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}

func (pr *ProductRepository) UpdateProduct(product models.Product) error {
	_, err := pr.Db.Exec("UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?",
		product.Name, product.Price, product.Stock, product.ID)
	return err
}

func (pr *ProductRepository) GetAllProducts() ([]models.Product, error) {
	rows, err := pr.Db.Query("SELECT id,name,price,stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) GetProductByName(name string) ([]models.Product, error) {
	rows, err := pr.Db.Query("SELECT id,name,price,stock FROM products WHERE name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) GetProductByID(id string) (models.Product, error) {
	row := pr.Db.QueryRow("SELECT id,name,price,stock FROM products WHERE id = ?", id)
	var product models.Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
