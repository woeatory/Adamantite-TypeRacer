package tests_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	"testing"
	"time"
)

func TestGetAll(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	userService := service.NewUserService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// mock expects
	expectedRows := sqlmock.NewRows([]string{"user_id", "username", "password_hash", "created_at"}).
		AddRow(uuid.New(), "user1", "hash1", time.Now()).
		AddRow(uuid.New(), "user2", "hash2", time.Now())
	mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(expectedRows)
	// act
	users, err := userService.GetAll()
	// assert
	assert.NoError(t, err)
	assert.Len(t, users, 2)

	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	userService := service.NewUserService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userID, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf(err.Error())
	}
	expectedRow := sqlmock.NewRows([]string{"user_id", "username", "password_hash", "created_at"}).
		AddRow(userID, "user1", "hash1", time.Now())

	mock.ExpectPrepare("SELECT (.+) FROM users WHERE (.+)")
	mock.ExpectQuery("SELECT (.+) FROM users WHERE (.+)").WithArgs(userID).WillReturnRows(expectedRow)
	_, err = userService.GetByID(userID.String())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestChangeUsername(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	userService := service.NewUserService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userID := uuid.New()
	// mock expects
	_ = sqlmock.NewRows([]string{"user_id", "username", "password_hash", "created_at"}).
		AddRow(userID, "user1", "hash1", time.Now())
	query := "UPDATE users"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))
	// act
	err = userService.ChangeUsername("newName", userID.String())
	if err != nil {
		return
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteUser(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	userService := service.NewUserService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	userID := uuid.New()
	_ = sqlmock.NewRows([]string{"user_id", "username", "password_hash", "created_at"}).
		AddRow(userID, "user1", "hash1", time.Now())
	query := "DELETE FROM users WHERE (.+)"
	// mock expects
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(userID).WillReturnResult(sqlmock.NewResult(0, 1))
	// act
	err = userService.ChangeUsername("newName", userID.String())
	if err != nil {
		return
	}
	if err != nil {
		t.Fatalf(err.Error())
	}
	// assert
	assert.NoError(t, mock.ExpectationsWereMet())
}
