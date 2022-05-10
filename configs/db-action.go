package configs

import (
	"context"
	"user_api/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

func migration(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), sql.Users)
	if err != nil {
		return err
	}

	_, err = db.Exec(context.Background(), sql.Func_create_validate)
	if err != nil {
		return err
	}

	return nil
}

func rollback(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), sql.R_users)
	if err != nil {
		return err
	}

	return nil
}
