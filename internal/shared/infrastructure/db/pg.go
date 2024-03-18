package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPgClient creates a new pgsql client
func NewPgClient(dbUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %+v\n", err)
	}

	if err := db.Ping(); err != nil {

		return nil, fmt.Errorf("failed to ping database: %+v\n", err)
	}

	return db, nil
}
