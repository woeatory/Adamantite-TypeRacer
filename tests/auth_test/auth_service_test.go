package auth_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/service"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestLogIn(t *testing.T) {
	db, mock, err := sqlmock.New()
	authService := service.NewAuthService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	hash, err := bcrypt.GenerateFromPassword([]byte("12345678"), bcrypt.DefaultCost)
	var user = DTO.UserDTO{
		Username: "user1",
		Password: "12345678",
	}
	row := sqlmock.NewRows([]string{"username", "password_hash"}).
		AddRow(user.Username, hash)
	query := "SELECT (.+) FROM users"
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(user.Username).WillReturnRows(row)
	_, err = authService.LogIn(user)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSignUp(t *testing.T) {
	db, mock, err := sqlmock.New()
	authService := service.NewAuthService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var newUser = DTO.UserDTO{
		Username: "user1",
		Password: "pass1",
	}
	query := "INSERT INTO users"
	mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
	_, err = authService.SignUp(newUser)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}
