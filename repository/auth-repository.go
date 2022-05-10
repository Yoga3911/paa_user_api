package repository

import (
	"context"
	"fmt"
	"user_api/models"
	"user_api/sql"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthR interface {
	InsertData(ctx context.Context, user models.Register, hash string) error
	VerifyData(ctx context.Context, email string) (models.User, error)
	CheckDuplicate(ctx context.Context, username string, email string) error
}

type authR struct {
	db *pgxpool.Pool
}

func NewAuthR(db *pgxpool.Pool) AuthR {
	return &authR{db: db}
}

func (a *authR) InsertData(ctx context.Context, user models.Register, hash string) error {
	_, err := a.db.Exec(ctx, sql.CreateUser, user.Username, user.Email, hash)

	return err
}

func (a *authR) VerifyData(ctx context.Context, email string) (models.User, error) {
	var usr models.User

	err := a.db.QueryRow(ctx, sql.VerifyCredential, email).Scan(&usr.ID, &usr.Username, &usr.Email, &usr.Password, &usr.CreateAt, &usr.UpdateAt)
	if err != nil {
		return usr, err
	}

	return usr, nil
}

func (a *authR) CheckDuplicate(ctx context.Context, username string, email string) error {
	var (
		usernameC, emailC int8
	)
	
	err := a.db.QueryRow(ctx, sql.RegisterVal, username, email).Scan(&usernameC, &emailC)
	if err != nil {
		return err
	}

	if usernameC != 0 {
		return fmt.Errorf("duplicate username")
	}

	if emailC != 0 {
		return fmt.Errorf("duplicate email")
	}

	return nil
}