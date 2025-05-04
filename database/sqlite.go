package database

import (
	"database/sql"
	"fmt"
)

func InitDb() (*sql.DB, error) {
	databaseFile := "test.db"

	db, err := sql.Open("sqlite3", databaseFile)

	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite db")

	}

	return db, nil
}
