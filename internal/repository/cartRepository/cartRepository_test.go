package cartRepository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/dto"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *CartRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	return db, mock, &CartRepository{db: db}
}

func TestCreateCart(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO cart (id, user_id) VALUES (?,?)")).
		WithArgs("cart1", "user1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.CreateCart("cart1", "user1"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAddToCart_NewItem(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	product := models.Product{ID: "p1"}

	// GetCartIDByUserID
	mock.ExpectQuery("SELECT id FROM cart").
		WithArgs("user1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("cart1"))

	// No existing product
	mock.ExpectQuery("SELECT product_id FROM cart_items").
		WithArgs("cart1", "p1").
		WillReturnError(sql.ErrNoRows)

	// Insert new
	mock.ExpectExec("INSERT INTO cart_items").
		WithArgs("cart1", "p1", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.AddToCart("user1", product); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAddToCart_ExistingItem(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	product := models.Product{ID: "p2"}

	// GetCartIDByUserID
	mock.ExpectQuery("SELECT id FROM cart").
		WithArgs("user2").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("cart2"))

	// Product already exists
	mock.ExpectQuery("SELECT product_id FROM cart_items").
		WithArgs("cart2", "p2").
		WillReturnRows(sqlmock.NewRows([]string{"product_id"}).AddRow("p2"))

	// Update quantity
	mock.ExpectExec("UPDATE cart_items SET quantity = quantity \\+ 1").
		WithArgs("p2").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.AddToCart("user2", product); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRemoveFromCart_QuantityMoreThanOne(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT quantity FROM cart_items").
		WithArgs("cart1", "p1").
		WillReturnRows(sqlmock.NewRows([]string{"quantity"}).AddRow(3))

	mock.ExpectExec("UPDATE cart_items SET quantity = quantity - 1").
		WithArgs("cart1", "p1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.RemoveFromCart("cart1", "p1"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRemoveFromCart_QuantityEqualOne(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT quantity FROM cart_items").
		WithArgs("cart1", "p1").
		WillReturnRows(sqlmock.NewRows([]string{"quantity"}).AddRow(1))

	mock.ExpectExec("DELETE FROM cart_items").
		WithArgs("cart1", "p1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.RemoveFromCart("cart1", "p1"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetCartIDByUserID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT id FROM cart").
		WithArgs("userX").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("cartX"))

	id, err := repo.GetCartIDByUserID("userX")
	if err != nil || id != "cartX" {
		t.Errorf("expected cartX got %s, err=%v", id, err)
	}
}

func TestEmptyCart(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	// GetCartIDByUserID
	mock.ExpectQuery("SELECT id FROM cart").
		WithArgs("userZ").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("cartZ"))

	// Delete items
	mock.ExpectExec("DELETE FROM cart_items").
		WithArgs("cartZ").
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.EmptyCart("userZ"); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetCartItemQuantity(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT quantity FROM cart_items").
		WithArgs("cartY", "prodY").
		WillReturnRows(sqlmock.NewRows([]string{"quantity"}).AddRow(4))

	q, err := repo.GetCartItemQuantity("cartY", "prodY")
	if err != nil || q != 4 {
		t.Errorf("expected 4 got %d, err=%v", q, err)
	}
}

func TestGetCartItems(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "quantity"}).
		AddRow("p1", "prod1", 100, 2).
		AddRow("p2", "prod2", 200, 1)

	mock.ExpectQuery("SELECT p.id, p.name, p.price, ci.quantity").
		WithArgs("cart123").
		WillReturnRows(rows)

	items, err := repo.GetCartItems("cart123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
	if items[0] != (dto.CartItemsDTO{ProductID: "p1", ProductName: "prod1", Price: 100, Quantity: 2}) {
		t.Errorf("unexpected item: %+v", items[0])
	}
}
