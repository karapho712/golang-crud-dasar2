package middleware

import (
	"crud-dasar-go-2/exception"
	"crud-dasar-go-2/helper"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("very-secret")

func CreateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		helper.PanicIfError(err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	if !token.Valid {
		panic(exception.NewUnauthorizedError("Invalid token"))
	}

	return nil
}
