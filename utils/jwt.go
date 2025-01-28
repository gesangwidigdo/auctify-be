package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": "auctify",
		"exp": time.Now().Add(time.Hour * 72).Unix(),
		"iat": time.Now().Unix(),
	})
	secretKey := os.Getenv("SECRET_KEY")
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
