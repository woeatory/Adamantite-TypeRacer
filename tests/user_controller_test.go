package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/controller"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUsersShouldReturnOk(t *testing.T) {
	// Arrange
	mockedUserID := uuid.New()
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockUsers := []model.User{
		{UserID: uuid.New(), Username: "user1"},
		{UserID: uuid.New(), Username: "user2"},
	}
	mockUserService.
		On("GetAll").
		Return(mockUsers, nil)

	router := setUpRouterWithAuth(mockedUserID.String())
	router.GET("/user/getAll", userController.GetAllUsers)

	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("GET", "/user/getAll", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"data": mockUsers}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockUserService.AssertCalled(t, "GetAll")
}

func TestGetAllUsersShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockUsers := []model.User{
		{UserID: uuid.New(), Username: "user1"},
		{UserID: uuid.New(), Username: "user2"},
	}
	mockUserService.
		On("GetAll").
		Return(mockUsers, nil)

	router := setUpRouterNotAuth()
	router.GET("/user/getAll", userController.GetAllUsers)

	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("GET", "/user/getAll", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockUserService.AssertNotCalled(t, "GetAll")
}

func TestGetAllUsersShouldReturnError(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	expectedErr := "unexpected error occurrence"
	mockUserService.
		On("GetAll").
		Return([]model.User{}, errors.New(expectedErr))

	router := setUpRouterWithAuth(uuid.New().String())
	router.GET("/user/getAll", userController.GetAllUsers)

	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("GET", "/user/getAll", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": expectedErr}
	// Assert the response
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockUserService.AssertCalled(t, "GetAll")
}

func TestGetUserByIdShouldReturnOk(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)

	// Define mock data
	mockedUserID := uuid.New()
	mockUser := model.User{UserID: mockedUserID, Username: "user1"}
	mockUserService.On("GetByID", mockedUserID.String()).Return(mockUser, nil)

	// Setup router
	router := setUpRouterWithAuth(mockedUserID.String())
	router.GET("/user/:userID", userController.GetUserByID)
	req, _ := http.NewRequest("GET", "/user/"+mockedUserID.String(), nil)
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	expectedResponse := gin.H{"data": mockUser}
	assert.JSONEq(t, rec.Body.String(), toJson(t, expectedResponse))

	// Assert the function calls
	mockUserService.AssertCalled(t, "GetByID", mockedUserID.String())
}

func TestGetUserByIdShouldReturnUnauthorized(t *testing.T) {
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)

	// Define mock data
	mockedUserID := uuid.New()
	mockUser := model.User{UserID: mockedUserID, Username: "user1"}
	mockUserService.On("GetByID", mockedUserID.String()).Return(mockUser, nil)

	// Setup router
	router := setUpRouterNotAuth()
	router.GET("/user/:userID", userController.GetUserByID)
	req, _ := http.NewRequest("GET", "/user/"+mockedUserID.String(), nil)
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	expectedResponse := gin.H{"error": "unauthorized"}
	assert.JSONEq(t, rec.Body.String(), toJson(t, expectedResponse))

	// Assert the function calls
	mockUserService.AssertNotCalled(t, "GetByID", mockedUserID.String())
}

func TestChangeUsernameShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockedUserID := uuid.New()
	mockNewUsername := "newUsername"

	mockInput := DTO.UserChangeUsernameDto{NewUsername: mockNewUsername}
	mockUserService.
		On("ChangeUsername", mockNewUsername, mockedUserID.String()).
		Return(nil)
	jsonValue, _ := json.Marshal(mockInput)
	// Setup router
	router := setUpRouterNotAuth()
	router.PUT("/user/changeUsername", userController.ChangeUsername)
	// Prepare a test request
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("PUT", "/user/changeUsername", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())

}

func TestChangeUsernameShouldReturnOk(t *testing.T) {
	// Arrange
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockedUserID := uuid.New()
	mockNewUsername := "newUsername"
	mockInput := DTO.UserChangeUsernameDto{NewUsername: mockNewUsername}
	mockUserService.
		On("ChangeUsername", mockNewUsername, mockedUserID.String()).
		Return(nil)
	jsonValue, _ := json.Marshal(mockInput)
	// Setup router
	router := setUpRouterWithAuth(mockedUserID.String())
	rec := httptest.NewRecorder()
	// Prepare a test request
	router.PUT("/user/changeUsername", userController.ChangeUsername)
	// Call the API handler
	req, _ := http.NewRequest("PUT", "/user/changeUsername", bytes.NewBuffer(jsonValue))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Successfully Changed Username"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert mock expectation
	mockUserService.AssertCalled(t, "ChangeUsername", mockNewUsername, mockedUserID.String())
}

func TestDeleteUserShouldReturnOk(t *testing.T) {
	// Arrange
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockedUserID := uuid.New()
	mockUserService.
		On("DeleteUser", mockedUserID.String()).
		Return(nil)
	// Setup router
	router := setUpRouterWithAuth(mockedUserID.String())
	rec := httptest.NewRecorder()
	// Prepare a test request
	router.DELETE("/user/deleteUser", userController.DeleteUser)
	// Call the API handler
	req, _ := http.NewRequest("DELETE", "/user/deleteUser", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Successfully Deleted User"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert mock expectation
	mockUserService.AssertCalled(t, "DeleteUser", mockedUserID.String())
}

func TestDeleteUserShouldReturnUnauthorized(t *testing.T) {
	// Arrange
	mockUserService := new(mocks.MockUserService)
	userController := controller.NewUserController(mockUserService)
	mockedUserID := uuid.New()
	mockUserService.
		On("DeleteUser", mockedUserID.String()).
		Return(nil)
	// Setup router
	router := setUpRouterNotAuth()
	rec := httptest.NewRecorder()
	// Prepare a test request
	router.DELETE("/user/deleteUser", userController.DeleteUser)
	// Call the API handler
	req, _ := http.NewRequest("DELETE", "/user/deleteUser", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": "unauthorized"}
	// Assert the response
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert mock expectation
	mockUserService.AssertNotCalled(t, "DeleteUser", mockedUserID.String())
}
