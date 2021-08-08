package adapter

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

var (
	db *sql.DB
)

func LoadMySQL() {
	db = MySQL()
}

// MySQL is open connection to mysql server
func MySQL() *sql.DB {
	db, _, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	return db
}

// DBSQL is open connection into database
func DBSQL() *sql.DB {
	return db
}
