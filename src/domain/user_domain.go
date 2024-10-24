package domain

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/lukinhas563/gochat/src/model/database/sqlite"
)

type UserDomain interface {
	CreateUser(request.UserRegister) error
	LoginUser(request.UserLogin) error
}

type userDomain struct {
	database sqlite.SqliteDatabase
}

func NewUserDomain(database sqlite.SqliteDatabase) UserDomain {
	return &userDomain{
		database: database,
	}
}

func (ud *userDomain) encrypt(password string) string {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) CreateUser(userRequest request.UserRegister) error {
	userRequest.Password = ud.encrypt(userRequest.Password)

	if err := ud.database.InsertUser(userRequest); err != nil {
		return err
	}

	return nil
}

func (ud *userDomain) LoginUser(userLogin request.UserLogin) error {

	user, err := ud.database.GetByUsername(userLogin.Username)
	if err != nil {
		return err
	}

	userLogin.Password = ud.encrypt(userLogin.Password)
	if user.Password != userLogin.Password {
		return fmt.Errorf("Invalid password")
	}

	return nil
}
