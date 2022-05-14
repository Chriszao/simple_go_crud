package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL Connection driver
)

// Open MySQL database connection
func Connect() (*sql.DB, error) {
	connectionString := "golang:golang@/go_course?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
