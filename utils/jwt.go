package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("TYHGFTYNVGHTJJHVCGHJNBVCGHJNBGHJNVGHJNBVBNNBV HNHB VBNBVB V")

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	return (*claims)["username"].(string), nil
}
