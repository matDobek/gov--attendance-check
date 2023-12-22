package storage

import (
	"database/sql"
	_ "github.com/lib/pq" // postgres driver for sql lib
)

type Storage struct {
	PrimaryDB *sql.DB
}

func NewStorage(databaseURL string) *Storage {
	db := NewSQLDatabase(databaseURL)

	return &Storage{
		PrimaryDB: db,
	}
}

func NewSQLDatabase(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		panic(err)
	}

	return db
}
