package service

import (
	"testing"

	"github.com/lukinhas563/gochat/src/model/api/response"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	tokenService := NewTokenService("secretkey")

	assert.NotNil(t, tokenService, "Expect tokenSerivce to not be nil")

	userLogin := response.UserLogin{
		Username: "testusername",
		Email:    "testemail@example.com",
		Password: "testpassword",
	}
	token, err := tokenService.GenerateToken(userLogin)

	assert.NoError(t, err, "Expect no error to generate a token")
	assert.NotNil(t, token, "Expect a token before to call GenerateToken method")
}
