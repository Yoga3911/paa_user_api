package repository

import (
	"context"
	"user_api/models"
	"user_api/sql"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserR interface {
	GetOne(ctx context.Context, id float64) pgx.Row
	Update(ctx context.Context, update models.Update, id float64) error
	CheckEmail(ctx context.Context, email string) pgx.Row
}

type userR struct {
	db *pgxpool.Pool
}

func NewUserR(db *pgxpool.Pool) UserR {
	return &userR{
		db: db,
	}
}

func (u *userR) GetOne(ctx context.Context, id float64) pgx.Row {
	pg := u.db.QueryRow(ctx, sql.GetOne, id)

	return pg
}

func (u *userR) Update(ctx context.Context, update models.Update, id float64) error {
	_, err := u.db.Exec(ctx, sql.UpdateUser, id, update.Username, update.Email, update.Image)

	return err
}

func (u *userR) CheckEmail(ctx context.Context, email string) pgx.Row {
	pg := u.db.QueryRow(ctx, sql.GetByEmail, email)

	return pg
}