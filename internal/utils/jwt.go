package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

/**
 * Claims are the data to be stored inside of the token
 * ID (userID), Email,
 * Permissions ( {[resourceId: int]: [permissionLevel: int]} map )
 */
type Claims struct {
	ID          int         `json:"id"`
	Email       string      `json:"email"`
	Permissions map[int]int `json:"permissions"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(userId int, email string, permissions map[int]int) (string, time.Time, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		ID:          userId,
		Email:       email,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	/*
		Do something like this on the handler
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	*/
	return tokenString, expirationTime, err
}
