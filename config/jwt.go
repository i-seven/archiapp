package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte
var JWTExpiry time.Duration

func init() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "dev-secret-change-me" // development fallback; override in prod
	}
	JWTSecret = []byte(secret)

	exp := os.Getenv("JWT_EXPIRY_MINUTES")
	if exp == "" {
		JWTExpiry = 60 * time.Minute
	} else {
		// parse minutes
		// ignore error for brevity; in production handle it
		d, _ := time.ParseDuration(exp + "m")
		JWTExpiry = d
	}
}

// CreateToken returns signed JWT for given user id and role.
func CreateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(JWTExpiry).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}
