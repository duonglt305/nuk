package db

import (
	"sync"

	"github.com/jmoiron/sqlx"
)

type IRds interface {
	Connect(dbUrl string) error
	Get() *sqlx.DB
}

const (
	PGSQLDriver = "pgsql"
	MYSQLDriver = "mysql"
)

var (
	dbIns IRds
	once  sync.Once
)

func New(driver string, dbUrl string) (IRds, error) {
	var err error
	once.Do(func() {
		switch driver {
		case PGSQLDriver:
			dbIns = &PostgresRds{}
		case MYSQLDriver:
			dbIns = &MySQLRds{}
		}
		err = dbIns.Connect(dbUrl)
	})
	return dbIns, err
}
