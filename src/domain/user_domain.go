package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/lukinhas563/gochat/src/model/database/sqlite"
	"github.com/lukinhas563/gochat/src/shared/service"
	"github.com/lukinhas563/gochat/src/shared/service/logger"
	"go.uber.org/zap"
)

type UserDomain interface {
	CreateUser(request.UserRegister) error
	LoginUser(request.UserLogin) (string, error)
}

type userDomain struct {
	database sqlite.SqliteDatabase
	token    service.TokenService
}

func NewUserDomain(database sqlite.SqliteDatabase, tokenService service.TokenService) UserDomain {
	return &userDomain{
		database: database,
		token:    tokenService,
	}
}

func (ud *userDomain) encrypt(password string) string {
	logger.Info("Init Encrypt from UserDomain", zap.String("journey", "encrypt"))

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))

	logger.Info("User's password encrypted successfully", zap.String("journey", "encrypt"))
	return hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) CreateUser(userRequest request.UserRegister) error {
	logger.Info("Init CreateUser from UserDomain", zap.String("journey", "CreateUser"))

	userRequest.Password = ud.encrypt(userRequest.Password)

	if err := ud.database.InsertUser(userRequest); err != nil {
		logger.Error("Error to inset user into database", err, zap.String("journey", "CreateUser"))

		return err
	}

	return nil
}

func (ud *userDomain) LoginUser(userLogin request.UserLogin) (string, error) {
	logger.Info("Init LoginUser from UserDomain", zap.String("journey", "LoginUser"))

	user, err := ud.database.GetByUsername(userLogin.Username)
	if err != nil {
		logger.Error("Error to found user into database", err, zap.String("journey", "LoginUser"))

		return "", err
	}

	userLogin.Password = ud.encrypt(userLogin.Password)
	if user.Password != userLogin.Password {
		logger.Error("Error to validate the user's password", err, zap.String("journey", "LoginUser"))

		return "", fmt.Errorf("Invalid password")
	}

	token, err := ud.token.GenerateToken(*user)
	if err != nil {
		logger.Error("Error to generate token", err, zap.String("journey", "LoginUser"))

		return "", err
	}

	return token, nil
}
