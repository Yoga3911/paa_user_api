package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTS interface {
	GenerateToken(id uint64, name string, email string, password string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtS struct {
	issuer    string
	secretKey string
}

type jwtCustomClaim struct {
	ID       uint64 `json:"id"`
	Username     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func NewJWTS() JWTS {
	return &jwtS{
		issuer:    "qwerty",
		secretKey: os.Getenv("JWT_SECRET"),
	}
}

func (j *jwtS) GenerateToken(id uint64, username string, email string, password string) string {
	claims := &jwtCustomClaim{
		id,
		username,
		email,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return string(err.Error())
	}

	return t
}

func (j *jwtS) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
}