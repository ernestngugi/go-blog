package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db DB

type DB interface {
	Close() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Ping() error
}

type ApplicationDB struct {
	*sql.DB
}

func InitDB() DB {
	return InitDBWithParams(
		os.Getenv("DATABASE_URL"),
	)
}

func InitDBWithParams(databaseURL string) DB {

	appDB := NewDatabaseWithURL(databaseURL)

	db = &ApplicationDB{
		DB: appDB,
	}

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func NewDatabaseWithURL(databaseURL string) *sql.DB {
	fmt.Println(databaseURL)
	db, err := sql.Open("mysql", databaseURL+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	return db
}
