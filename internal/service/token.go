package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var signingKey = os.Getenv("SIGN_KEY_STRING")

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

func CreateToken(userId string) (string, error) {
	if userId == "" {
		return "", errors.New("user_id cannot be empty")
	}

	params := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		UserId: userId,
	})

	token, err := params.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	if claims.UserId == "" {
		return "", errors.New("user_id is empty or invalid")
	}

	return claims.UserId, nil
}
