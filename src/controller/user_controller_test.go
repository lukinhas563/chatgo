package controller

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

type mockDomain struct {
	mock.Mock
}

func (m *mockDomain) CreateUser(userRequest request.UserRegister) error {
	args := m.Called(userRequest)
	return args.Error(0)
}

func (m *mockDomain) LoginUser(userLogin request.UserLogin) (string, error) {
	args := m.Called(userLogin)
	return args.String(0), args.Error(1)
}

func TestUserController_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDomain := new(mockDomain)
	userController := NewUserController(mockDomain)

	registerRequest := request.UserRegister{
		Username: "testusername",
		Email:    "testemail@example.com",
		Password: "testpassword",
	}

	mockDomain.On("CreateUser", registerRequest).Return(nil)

	// Codifica o registerRequest em JSON
	reqBody, _ := json.Marshal(registerRequest)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Cria um response recorder para capturar a resposta do controlador
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	userController.Register(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, `"Registered successfully"`, rec.Body.String())
	mockDomain.AssertExpectations(t)
}

func TestUserController_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDomain := new(mockDomain)
	userController := NewUserController(mockDomain)

	loginRequest := request.UserLogin{
		Username: "testusername",
		Password: "testpassword",
	}

	mockDomain.On("LoginUser", loginRequest).Return("mockToken", nil)

	// Codifica o loginRequest em JSON
	reqBody, _ := json.Marshal(loginRequest)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Cria um response recorder para capturar a resposta do controlador
	rec := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	userController.Login(ctx)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"mockToken"`, rec.Body.String())
	mockDomain.AssertExpectations(t)
}
