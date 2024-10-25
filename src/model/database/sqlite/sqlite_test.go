package sqlite

import (
	"fmt"
	"testing"

	"github.com/lukinhas563/gochat/src/model/api/request"
	"github.com/stretchr/testify/assert"
)

func TestSqliteDatabase_Connect(t *testing.T) {
	database_name := ":memory:"

	sqliteDB := NewSqliteDatabase()
	err := sqliteDB.Connect(database_name)
	defer sqliteDB.Close()

	assert.NoError(t, err, "Expected no error when connecting to in-memory database")
}

func TestSqliteDatabase_Connect_error(t *testing.T) {
	database_name := ""

	sqliteDB := NewSqliteDatabase()
	err := sqliteDB.Connect(database_name)
	defer sqliteDB.Close()

	fmt.Println(err)
	assert.Error(t, err, "Expected an error when connecting to a invalible database file")
}

func TestSqliteDatabase_InsertUser(t *testing.T) {
	database_name := ":memory:"

	sqliteDB := NewSqliteDatabase()
	err := sqliteDB.Connect(database_name)
	assert.NoError(t, err, "Expected no error when connecting to in-memory database")
	defer sqliteDB.Close()

	_, err = sqliteDB.(*sqliteDatabase).database.Exec(`
			CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`)
	assert.NoError(t, err, "Expected no error when creating table")

	user := request.UserRegister{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword123",
	}

	err = sqliteDB.InsertUser(user)
	assert.NoError(t, err, "Expected no error when inserting user")
}

func TestSqliteDatabase_InsertUser_error(t *testing.T) {
	sqliteDB := NewSqliteDatabase()
	defer sqliteDB.Close()

	user := request.UserRegister{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword123",
	}

	err := sqliteDB.InsertUser(user)

	assert.Error(t, err, "Expected an error when inserting user into a database not connected")
}

func TestSqliteDatabase_GetByUsername(t *testing.T) {
	database_name := ":memory:"

	sqliteDB := NewSqliteDatabase()
	err := sqliteDB.Connect(database_name)
	assert.NoError(t, err, "Expected no error when connecting to in-memory database")
	defer sqliteDB.Close()

	_, err = sqliteDB.(*sqliteDatabase).database.Exec(`
			CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`)
	assert.NoError(t, err, "Expected no error when creating table")

	user := request.UserRegister{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword123",
	}

	err = sqliteDB.InsertUser(user)
	assert.NoError(t, err, "Expected no error when inserting user")

	userData, err := sqliteDB.GetByUsername("testuser")
	assert.NoError(t, err, "Expected no error when retrieving user by username")
	assert.NotNil(t, userData, "Expected user data to be returned")
	assert.Equal(t, "testuser", userData.Username, "Expected correct username")
	assert.Equal(t, "testuser@example.com", userData.Email, "Expected correct email")
}

func TestSqliteDatabase_GetByUsername_Error(t *testing.T) {
	database_name := ":memory:"

	sqliteDB := NewSqliteDatabase()
	err := sqliteDB.Connect(database_name)
	assert.NoError(t, err, "Expected no error when connecting to in-memory database")
	defer sqliteDB.Close()

	_, err = sqliteDB.(*sqliteDatabase).database.Exec(`
			CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`)
	assert.NoError(t, err, "Expected no error when creating table")

	user := request.UserRegister{
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword123",
	}

	err = sqliteDB.InsertUser(user)
	assert.NoError(t, err, "Expected no error when inserting user")

	userData, err := sqliteDB.GetByUsername("testuser2")
	assert.Error(t, err, "Expected an error when when retrieving user by invalid username")
	assert.Nil(t, userData, "Expected a nil when when retrieving user by invalid username")
}
