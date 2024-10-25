package domain

import (
	"testing"

	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/lukinhas563/gochat/src/model/api/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDatabase struct {
	mock.Mock
}

func (m *mockDatabase) Connect(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *mockDatabase) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *mockDatabase) InsertUser(user request.UserRegister) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockDatabase) GetByUsername(username string) (*response.UserLogin, error) {
	args := m.Called(username)
	return args.Get(0).(*response.UserLogin), args.Error(1)
}

type mockTokenService struct {
	mock.Mock
}

func (m *mockTokenService) GenerateToken(user response.UserLogin) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

type mockEmailService struct {
	mock.Mock
}

func (m *mockEmailService) Send(recipient, message string) error {
	args := m.Called(recipient, message)
	return args.Error(0)
}

func TestUserDomain_CreateUser(t *testing.T) {
	mockDB := new(mockDatabase)
	mockToken := new(mockTokenService)
	mockEmail := new(mockEmailService)
	userDomain := NewUserDomain(mockDB, mockToken, mockEmail)

	user := request.UserRegister{
		Username: "testusername",
		Email:    "testemail@example.com",
		Password: "testpassword",
	}

	mockDB.On("InsertUser", mock.AnythingOfType("request.UserRegister")).Return(nil)

	err := userDomain.CreateUser(user)

	assert.NoError(t, err, "Expect no error to create a new user")
	mockDB.AssertExpectations(t)
}

func TestUserDOmain_LoginUser(t *testing.T) {
	mockDB := new(mockDatabase)
	mockToken := new(mockTokenService)
	mockEmail := new(mockEmailService)
	domainService := NewUserDomain(mockDB, mockToken, mockEmail)

	user := request.UserLogin{
		Username: "testusername",
		Password: "testpassword",
	}

	userData := &response.UserLogin{
		Username: "testusername",
		Email:    "testemail@example.com",
		Password: domainService.(*userDomain).encrypt("testpassword"),
	}

	mockDB.On("GetByUsername", "testusername").Return(userData, nil)
	mockToken.On("GenerateToken", *userData).Return("mockToken", nil)

	token, err := domainService.LoginUser(user)

	assert.NoError(t, err, "Expect no error to create a new user")
	assert.Equal(t, "mockToken", token, "Expect the token to be the same at the start")
	mockDB.AssertExpectations(t)
	mockToken.AssertExpectations(t)
}
