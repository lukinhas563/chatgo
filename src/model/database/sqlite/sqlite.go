package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/lukinhas563/gochat/src/model/api/request"
	_ "modernc.org/sqlite"
)

type SqliteDatabase interface {
	Connect(string) error
	Close() error
	InsertUser(request.UserRegister) error
}
type sqliteDatabase struct {
	database *sql.DB
}

func NewSqliteDatabase() SqliteDatabase {
	return &sqliteDatabase{}
}

func (sqlite *sqliteDatabase) Connect(path string) error {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	sqlite.database = db
	return nil
}

func (sqlite *sqliteDatabase) Close() error {
	if sqlite.database != nil {
		return sqlite.database.Close()
	}
	return nil
}

func (sqlite *sqliteDatabase) InsertUser(user request.UserRegister) error {
	if sqlite.database != nil {

		query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
		_, err := sqlite.database.Exec(query, user.Username, user.Email, user.Password)

		if err != nil {
			return err
		}

		fmt.Println("User inserted successfully")
		return nil
	}

	return fmt.Errorf("Database not connected")
}
