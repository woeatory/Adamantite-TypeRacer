package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/controller"
	"github.com/woeatory/Adamantite-TypeRacer/internal/auth/models/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/tests/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUpOk(t *testing.T) {
	mockAuthService := new(mocks.MockAuthService)
	authController := controller.NewAuthController(mockAuthService)
	mockAuthService.
		On("SignUp", mock.AnythingOfType("DTO.UserDTO")).
		Return("uuid", nil)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.POST("/auth/signup", authController.SignUp)

	// Prepare a test request
	var input = DTO.UserDTO{
		Username: "username1",
		Password: "password1",
	}
	payload, _ := json.Marshal(input)

	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Successfully Signed Up"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockAuthService.AssertCalled(t, "SignUp", mock.AnythingOfType("DTO.UserDTO"))
}

func TestSignUpDuplicate(t *testing.T) {
	mockAuthService := new(mocks.MockAuthService)
	authController := controller.NewAuthController(mockAuthService)
	duplicateErr := "duplicate key value violates unique constraint \"users_username_key\""
	mockAuthService.
		On("SignUp", mock.AnythingOfType("DTO.UserDTO")).
		Return("", errors.New(duplicateErr))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.POST("/auth/signup", authController.SignUp)

	// Prepare a test request
	var input = DTO.UserDTO{
		Username: "username1",
		Password: "password1",
	}
	payload, _ := json.Marshal(input)

	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": duplicateErr}
	// Assert the response
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockAuthService.AssertCalled(t, "SignUp", mock.AnythingOfType("DTO.UserDTO"))
}

func TestLogInOk(t *testing.T) {
	mockAuthService := new(mocks.MockAuthService)
	authController := controller.NewAuthController(mockAuthService)
	mockAuthService.
		On("LogIn", mock.AnythingOfType("DTO.UserDTO")).
		Return("uudi", nil)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.POST("/auth/login", authController.LogIn)

	// Prepare a test request
	var input = DTO.UserDTO{
		Username: "username1",
		Password: "password1",
	}
	payload, _ := json.Marshal(input)

	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Successfully Logged In"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockAuthService.AssertCalled(t, "LogIn", mock.AnythingOfType("DTO.UserDTO"))
}

func TestLogInBadRequest(t *testing.T) {
	mockAuthService := new(mocks.MockAuthService)
	authController := controller.NewAuthController(mockAuthService)
	errorMessage := "wrong"
	mockAuthService.
		On("LogIn", mock.AnythingOfType("DTO.UserDTO")).
		Return("", errors.New(errorMessage))
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	router.POST("/auth/login", authController.LogIn)

	// Prepare a test request
	var input = DTO.UserDTO{
		Username: "username1",
		Password: "password1",
	}
	payload, _ := json.Marshal(input)

	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(payload))
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"error": errorMessage}
	// Assert the response
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())
	// Assert the function calls
	mockAuthService.AssertCalled(t, "LogIn", mock.AnythingOfType("DTO.UserDTO"))
}

func TestLogOutShouldReturnOk(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("testsesesion", store))
	mockAuthService := new(mocks.MockAuthService)
	authController := controller.NewAuthController(mockAuthService)
	router.POST("/auth/logout", authController.LogOut)
	rec := httptest.NewRecorder()
	// Call the API handler
	req, _ := http.NewRequest("POST", "/auth/logout", nil)
	router.ServeHTTP(rec, req)
	expectedResponse := gin.H{"message": "Successfully Logged Out"}
	// Assert the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, toJson(t, expectedResponse), rec.Body.String())

}

func toJson(t *testing.T, v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return string(jsonStr)
}
