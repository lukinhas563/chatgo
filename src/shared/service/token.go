package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lukinhas563/gochat/src/model/api/response"
	resterr "github.com/lukinhas563/gochat/src/shared/service/restErr"
)

type TokenService interface {
	GenerateToken(response.UserLogin) (string, error)
}

type tokenService struct {
	secret string
}

func NewTokenService(secret string) TokenService {
	return &tokenService{
		secret: secret,
	}
}

func (ts *tokenService) GenerateToken(user response.UserLogin) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(ts.secret))
	if err != nil {
		return "", resterr.NewInternalServerError("Error trying to generate JWT token")
	}
	return tokenString, nil
}
