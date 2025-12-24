package config

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("RAHASIA_SUPER") // nanti bisa pakai env

func GenerateToken(userID uint, username string, role string, nama_lengkap string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role":     role,
		"nama_lengkap":     nama_lengkap,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 1 hari
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
