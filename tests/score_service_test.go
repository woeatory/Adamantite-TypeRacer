package tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/woeatory/Adamantite-TypeRacer/internal/repository"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/service"
	"testing"
)

func TestNewRecord(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	scoreService := service.NewScoreService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var wpm, accuracy, typos = 1, 1, 0
	userID := uuid.New()
	// mock expects
	query := "INSERT INTO scores"
	sqlmock.NewRows([]string{"userID", "WPM", "accuracy", "typos"})
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(userID, wpm, accuracy, typos).WillReturnResult(sqlmock.NewResult(0, 1))
	// act
	err = scoreService.NewScoreRecord(userID.String(), wpm, accuracy, typos)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetRecordsByUserID(t *testing.T) {
	// arrange
	recordID, userID := 1, uuid.New()
	db, mock, err := sqlmock.New()
	scoreService := service.NewScoreService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// mock expects
	expected := mock.NewRows([]string{"record_id", "user_id", "wpm", "accuracy", "typos"}).
		AddRow(recordID, userID.String(), 1, 1, 1)
	//AddRow(2, uuid.New().String())
	query := "SELECT (.+) FROM scores WHERE (.+)"
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(userID).WillReturnRows(expected)
	// act
	_, err = scoreService.GetRecordsByUserID(userID.String())
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteRecord(t *testing.T) {
	// arrange
	recordID, userID := 1, uuid.New()
	db, mock, err := sqlmock.New()
	scoreService := service.NewScoreService(&repository.Repo{DB: db})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// mock expects
	mock.NewRows([]string{"record_id", "user_id"}).
		AddRow(recordID, userID.String())
	query := "DELETE FROM scores WHERE (.+)"
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(recordID, userID).WillReturnResult(sqlmock.NewResult(0, 1))
	// act
	err = scoreService.DeleteScoreRecord(userID.String(), recordID)
	if err != nil {
		t.Fatalf(err.Error())
	}
	// Assert that the mock query was executed as expected
	assert.NoError(t, mock.ExpectationsWereMet())
}
