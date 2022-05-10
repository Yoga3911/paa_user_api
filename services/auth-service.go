package services

import (
	"context"
	"fmt"
	"strings"
	"user_api/models"
	"user_api/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthS interface {
	CreateUser(ctx context.Context, user models.Register) error
	VerifyCredential(ctx context.Context, user models.Login) (string, models.User, error)
}

type authS struct {
	authR repository.AuthR
	jwtS  JWTS
}

func NewAuthS(authR repository.AuthR, jwtS JWTS) AuthS {
	return &authS{authR: authR, jwtS: jwtS}
}

func (a *authS) CreateUser(ctx context.Context, user models.Register) error {
	err := a.authR.CheckDuplicate(ctx, user.Username, user.Email)
	if err != nil {
		return err
	}	

	hash, err := hashAndSalt(user.Password)
	if err != nil {
		return err
	}

	err = a.authR.InsertData(ctx, user, hash)
	if err != nil {
		return err
	}

	return nil
}

func (a *authS) VerifyCredential(ctx context.Context, user models.Login) (string, models.User, error) {
	usr, err := a.authR.VerifyData(ctx, user.Email)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return "", usr, fmt.Errorf("email not found")
		}
		return "", usr, fmt.Errorf(err.Error())
	}

	err = comparePwd([]byte(usr.Password), user.Password)
	if err != nil {
		return "", usr, fmt.Errorf("wrong password")
	}

	token := a.jwtS.GenerateToken(usr.ID, usr.Username, usr.Email, usr.Password)

	return token, usr, nil
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func comparePwd(hash []byte, pwd string) error {
	err := bcrypt.CompareHashAndPassword(hash, []byte(pwd))

	return err
}