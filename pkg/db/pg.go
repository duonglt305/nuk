package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRds struct {
	db *sqlx.DB
}

func (rds *PostgresRds) Connect(dbUrl string) error {
	var err error
	rds.db, err = sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %+v\n", err)
	}
	if err := rds.db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %+v\n", err)
	}
	return nil
}

func (rds *PostgresRds) Get() *sqlx.DB {
	return rds.db
}
