package repository

import (
	"context"
	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/000001_init_schema.up.sql
var initialSchema string

func RunInitialSchema(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, initialSchema)
	return err
}
