package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPgClient creates a new pgsql client
func NewPgClient(dbUrl string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		fmt.Printf("failed to connect to database: %+v\n", err)
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("failed to ping database: %+v\n", err)
		return nil
	}

	return db
}
