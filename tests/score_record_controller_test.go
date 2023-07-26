package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/controller"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user_scores/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddNewScoreRecordShouldReturnOk(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	wpm, accuracy, typos := 1, 1, 1

	mockScoreService.
		On("NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos).
		Return(nil)

	router := setUpRouterWithAuth(mockedUserID.String())
	router.POST("/score/newScoreRecord", scoreController.AddNewScoreRecord)
	var input = DTO.NewScoreRecordDTO{
		WPM:      wpm,
		Accuracy: accuracy,
		Typos:    typos,
	}
	payload, _ := json.Marshal(input)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/score/newScoreRecord", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Record Added Successfully"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertCalled(t, "NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos)
}

func TestAddNewScoreRecordShouldReturnError(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	mockedError := errors.New("error text")
	wpm, accuracy, typos := 1, 1, 1
	mockScoreService.
		On("NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos).
		Return(mockedError)

	router := setUpRouterWithAuth(mockedUserID.String())
	router.POST("/score/newScoreRecord", scoreController.AddNewScoreRecord)
	var input = DTO.NewScoreRecordDTO{
		WPM:      wpm,
		Accuracy: accuracy,
		Typos:    typos,
	}
	payload, _ := json.Marshal(input)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/score/newScoreRecord", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "error while inserting new record:" + mockedError.Error()}
	// Assert the response
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertCalled(t, "NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos)
}

func TestAddNewScoreRecordShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	wpm, accuracy, typos := 1, 1, 1
	mockScoreService.
		On("NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos).
		Return(nil)

	router := setUpRouterNotAuth()
	router.POST("/score/newScoreRecord", scoreController.AddNewScoreRecord)
	var input = DTO.NewScoreRecordDTO{
		WPM:      wpm,
		Accuracy: accuracy,
		Typos:    typos,
	}
	payload, _ := json.Marshal(input)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/score/newScoreRecord", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertNotCalled(t, "NewScoreRecord", mockedUserID.String(), wpm, accuracy, typos)
}

func TestDeleteScoreRecordShouldReturnOk(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	mockedRecordId := 1
	mockScoreService.
		On("DeleteScoreRecord", mockedUserID.String(), mockedRecordId).
		Return(nil)

	router := setUpRouterWithAuth(mockedUserID.String())
	router.DELETE("/score/deleteScoreRecord", scoreController.DeleteScoreRecord)
	var input = DTO.DeleteScoreRecord{RecordID: mockedRecordId}
	payload, _ := json.Marshal(input)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("DELETE", "/score/deleteScoreRecord", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Record Deleted Successfully"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertCalled(t, "DeleteScoreRecord", mockedUserID.String(), mockedRecordId)
}

func TestDeleteScoreRecordShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	mockedRecordId := 1
	mockScoreService.
		On("DeleteScoreRecord", mockedUserID.String(), mockedRecordId).
		Return(nil)

	router := setUpRouterNotAuth()
	router.DELETE("/score/deleteScoreRecord", scoreController.DeleteScoreRecord)
	var input = DTO.DeleteScoreRecord{RecordID: mockedRecordId}
	payload, _ := json.Marshal(input)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("DELETE", "/score/deleteScoreRecord", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertNotCalled(t, "DeleteScoreRecord", mockedUserID.String(), mockedRecordId)
}

func TestGetAllUsersRecordsShouldReturnOkAndArrayOfRecords(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)
	mockedResult := []models.ScoreRecord{
		{
			RecordID:  1,
			UserID:    mockedUserID,
			WPM:       60,
			Accuracy:  90,
			Typos:     2,
			CreatedAt: time.Now(),
		},
		{
			RecordID:  2,
			UserID:    mockedUserID,
			WPM:       70,
			Accuracy:  95,
			Typos:     1,
			CreatedAt: time.Now(),
		},
	}
	mockScoreService.
		On("GetAllUsersScoreRecords", mockedUserID.String()).
		Return(mockedResult, nil)

	router := setUpRouterWithAuth(mockedUserID.String())
	router.GET("/score/getAllUsersScoreRecords", scoreController.GetAllUsersRecords)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("GET", "/score/getAllUsersScoreRecords", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": mockedResult}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertCalled(t, "GetAllUsersScoreRecords", mockedUserID.String())
}

func TestGetAllUsersRecordsShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockScoreService := new(mocks.MockScoreRecord)
	scoreController := controller.NewScoreController(mockScoreService)

	router := setUpRouterNotAuth()
	router.GET("/score/getAllUsersScoreRecords", scoreController.GetAllUsersRecords)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("GET", "/score/getAllUsersScoreRecords", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockScoreService.AssertNotCalled(t, "GetAllUsersScoreRecords", mockedUserID.String())
}
