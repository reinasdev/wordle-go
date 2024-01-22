package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
	"wordle/config"
)

// Database is the database connection
type Database struct {
	mu sync.Mutex
	DB *sql.DB
}

func LoadDatabase() (*Database, error) {
	cfg := config.GetConfig()

	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}
