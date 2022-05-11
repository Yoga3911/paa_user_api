package services

import (
	"context"
	"fmt"
	"user_api/models"
	"user_api/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserS interface {
	GetOne(ctx context.Context, token string) (models.User, error)
	Update(ctx context.Context, update models.Update, token string) (string, error)
	CheckEmail(ctx context.Context, email string, token string) error
}

type userS struct {
	userR repository.UserR
	db    *pgxpool.Pool
	jwtS  JWTS
}

func NewUserS(db *pgxpool.Pool, userR repository.UserR, jwtS JWTS,) UserS {
	return &userS{db: db, userR: userR, jwtS: jwtS}
}

func (u *userS) GetOne(ctx context.Context, token string) (models.User, error) {
	var user models.User

	t, err := u.jwtS.ValidateToken(token)
	if err != nil {
		return user, err
	}

	claims := t.Claims.(jwt.MapClaims)

	err = u.userR.GetOne(ctx, claims["id"].(float64)).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Image, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userS) Update(ctx context.Context, update models.Update, t string) (string, error) {
	valToken, err := u.jwtS.ValidateToken(t)
	if err != nil {
		return "", err
	}

	claims := valToken.Claims.(jwt.MapClaims)

	err = u.userR.Update(ctx, update, claims["id"].(float64))
	if err != nil {
		return "", err
	}

	token := u.jwtS.GenerateToken(uint64(claims["id"].(float64)), update.Username, update.Email, claims["password"].(string))

	return token, nil
}

func (u *userS) CheckEmail(ctx context.Context, email string, token string) error {
	var count int
	u.userR.CheckEmail(ctx, email).Scan(&count)
	t, err := u.jwtS.ValidateToken(token)
	if err != nil {
		return err
	}

	claims := t.Claims.(jwt.MapClaims)

	if count > 0 && email != claims["email"] {
		return fmt.Errorf("duplicate email")
	}

	return nil
}