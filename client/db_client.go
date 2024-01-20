package client

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type (
	Post struct {
		UserId int    `json:"userId" db:"userId"`
		Id     int    `json:"id" db:"id"`
		Title  string `json:"title" db:"title"`
		Body   string `json:"body" db:"body"`
	}

	Comment struct {
		PostId int    `json:"postId" db:"postId"`
		Id     int    `json:"id" db:"id"`
		Name   string `json:"name" db:"name"`
		Email  string `json:"email" db:"email"`
		Body   string `json:"body" db:"body"`
	}
)

var DBClient *sql.DB

func InitConnection() {
	db, err := sql.Open("mysql", "DenysShpak:den132435123shpkQWE@/media_db")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		return
	}

	DBClient = db
}
