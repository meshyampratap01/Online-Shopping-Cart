package productRepository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, ProductManager) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	return db, mock, &ProductRepository{Db: db}
}

func TestAddProduct(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("INSERT INTO products").
		WithArgs("1", "Product1", 100.0, 10).
		WillReturnResult(sqlmock.NewResult(1, 1))

	product := models.Product{ID: "1", Name: "Product1", Price: 100.0, Stock: 10}
	if err := repo.AddProduct(product); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}


func TestRemoveProduct(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("DELETE FROM products").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.RemoveProduct("1"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestUpdateProduct(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec("UPDATE products").
		WithArgs("UpdatedProduct", 150.0, 20, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	product := models.Product{ID: "1", Name: "UpdatedProduct", Price: 150.0, Stock: 20}
	if err := repo.UpdateProduct(product); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetAllProducts(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM products").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
			AddRow("1", "Product1", 100.0, 10).
			AddRow("2", "Product2", 200.0, 20))

	products, err := repo.GetAllProducts()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []models.Product{
		{ID: "1", Name: "Product1", Price: 100.0, Stock: 10},
		{ID: "2", Name: "Product2", Price: 200.0, Stock: 20},
	}

	if len(products) != len(expected) {
		t.Errorf("expected %d products, got %d", len(expected), len(products))
	}
}

func TestGetProductByName(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM products WHERE name LIKE ?").
		WithArgs("%Product%").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
			AddRow("1", "Product1", 100.0, 10).
			AddRow("2", "Product2", 200.0, 20))

	name := "Product"
	products, err := repo.GetProductByName(&name)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := []models.Product{
		{ID: "1", Name: "Product1", Price: 100.0, Stock: 10},
		{ID: "2", Name: "Product2", Price: 200.0, Stock: 20},
	}

	if len(products) != len(expected) {
		t.Errorf("expected %d products, got %d", len(expected), len(products))
	}
}


func TestGetProductByID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM products WHERE id = ?").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "stock"}).
			AddRow("1", "Product1", 100.0, 10))

	product, err := repo.GetProductByID("1")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expected := models.Product{ID: "1", Name: "Product1", Price: 100.0, Stock: 10}
	if product != expected {
		t.Errorf("expected %+v, got %+v", expected, product)
	}
}
