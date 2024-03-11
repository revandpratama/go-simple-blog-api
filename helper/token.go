package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/revandpratama/go-simple-blog-api/entity"
)

type JWTCustomClaims struct {
	ID int
	jwt.RegisteredClaims
}

var privateKey []byte

func GenerateToken(user *entity.User) (string, error) {
	claims := &JWTCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)), // expire after
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	ss, err := token.SignedString(privateKey)

	return ss, err
}

func ValidateToken(tokenString string) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenString, JWTCustomClaims{}, func(t *jwt.Token) (any, error) {
		return privateKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("token invalid")
		} else {
			return nil, errors.New("token expired")
		}
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token expired")
	}

	return &claims.ID, nil
}
