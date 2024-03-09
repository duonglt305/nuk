package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// NewPgClient creates a new PostgreSQL client
func NewPgClient(dbUrl string) *pgx.Conn {
	ctx := context.Background()
	db, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		fmt.Printf("failed to connect to database: %+v\n", err)
		return nil
	}

	if err := db.Ping(ctx); err != nil {
		fmt.Printf("failed to ping database: %+v\n", err)
		return nil
	}

	return db
}
