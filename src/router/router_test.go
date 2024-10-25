package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserController struct {
	mock.Mock
}

func (m *mockUserController) Register(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, "Registered successfully")
}

func (m *mockUserController) Login(c *gin.Context) {
	m.Called(c)
	c.JSON(http.StatusOK, "Login successful")
}

func (m *mockUserController) Confirm(c *gin.Context) {
	c.JSON(http.StatusOK, "User Confirmed")
}

func (m *mockUserController) Send(c *gin.Context) {
	c.JSON(http.StatusOK, "User Send reset")
}

func (m *mockUserController) Reset(c *gin.Context) {
	c.JSON(http.StatusOK, "User reset password")
}

func TestRouter_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	apiGroup := router.Group("/api")
	mockController := new(mockUserController)
	InitRouter(apiGroup, mockController)

	mockController.On("Register", mock.Anything).Return()

	registerReqiest := request.UserRegister{
		Username: "testusername",
		Email:    "testemail@example.com",
		Password: "testpassword",
	}

	reqBody, _ := json.Marshal(registerReqiest)
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Cria um response recorder para capturar a resposta do controlador
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `"Registered successfully"`, rec.Body.String())
	mockController.AssertExpectations(t)
}

func TestRouter_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	apiGroup := router.Group("/api")
	mockController := new(mockUserController)
	InitRouter(apiGroup, mockController)

	mockController.On("Login", mock.Anything).Return()

	loginRequest := request.UserLogin{
		Username: "testusername",
		Password: "testpassword",
	}

	reqBody, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `"Login successful"`, rec.Body.String())
	mockController.AssertExpectations(t)
}
