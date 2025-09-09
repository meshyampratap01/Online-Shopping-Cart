package userRepository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/meshyampratap01/OnlineShoppingCart/internal/models"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, UserManager) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	return db, mock, &UserRepository{Db: db}
}

func TestSaveUser(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (id, name, email, password, role) VALUES (?, ?, ?, ?, ?)")).
		WithArgs("1", "John Doe", "john@example.com", "password123", models.Customer).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := models.User{ID: "1", Name: "John Doe", Email: "john@example.com", Password: "password123", Role: models.Customer}
	if err := repo.SaveUser(user); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetUserByID(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, email, password, role FROM users").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role"}).
			AddRow("1", "John Doe", "john@example.com", "password123", models.Customer))

	user, err := repo.GetUserByID("1")
	if err != nil || user.ID != "1" || user.Name != "John Doe" || user.Email != "john@example.com" || user.Password != "password123" || user.Role != models.Customer {
		t.Errorf("unexpected user: %+v, err: %v", user, err)
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, email, password, role FROM users").
		WithArgs("john@example.com").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "role"}).
			AddRow("1", "John Doe", "john@example.com", "password123", models.Customer))

	user, err := repo.GetUserByEmail("john@example.com")
	if err != nil || user.ID != "1" || user.Name != "John Doe" || user.Email != "john@example.com" || user.Password != "password123" || user.Role != models.Customer {
		t.Errorf("unexpected user: %+v, err: %v", user, err)
	}
}
