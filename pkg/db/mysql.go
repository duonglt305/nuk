package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQLRds struct {
	db *sqlx.DB
}

func (rds *MySQLRds) Connect(dbUrl string) error {
	var err error
	rds.db, err = sqlx.Connect("mysql", dbUrl)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %+v\n", err)
	}
	if err := rds.db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %+v\n", err)
	}
	return nil
}

func (rds *MySQLRds) Get() *sqlx.DB {
	return rds.db
}
