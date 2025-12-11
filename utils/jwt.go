package utils

import (
	"errors"
	"fmt"
	"net/http"

	"backendAf/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a signed JWT for the given user ID and role.
// It delegates expiry/secret configuration to config.CreateToken.
func GenerateToken(userID uint, role string) (string, error) {
	return config.CreateToken(userID, role)
}

// ParseToken validates and parses a JWT string returning the token and claims map.
// It verifies signature with the secret in config.JWTSecret.
func ParseToken(tokenStr string) (*jwt.Token, jwt.MapClaims, error) {
	if tokenStr == "" {
		return nil, nil, errors.New("token is empty")
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		// enforce HMAC signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return token, nil, errors.New("invalid token claims")
	}
	return token, claims, nil
}

// GetUserIDAndRoleFromToken parses tokenStr and extracts the "sub" (user id) and "role" claims.
// Returns userID (uint), role (string) and error if parsing or claim types are invalid.
func GetUserIDAndRoleFromToken(tokenStr string) (uint, string, error) {
	_, claims, err := ParseToken(tokenStr)
	if err != nil {
		return 0, "", err
	}

	sub, ok := claims["sub"]
	if !ok {
		return 0, "", errors.New("sub claim missing")
	}

	var uid uint
	switch v := sub.(type) {
	case float64:
		uid = uint(v)
	case int:
		uid = uint(v)
	case uint:
		uid = v
	default:
		return 0, "", errors.New("invalid sub claim type")
	}

	role := ""
	if r, ok := claims["role"]; ok {
		if rs, ok := r.(string); ok {
			role = rs
		}
	}

	return uid, role, nil
}

func JSONError(c *gin.Context, status int, code string, message string) {
	c.JSON(status, gin.H{
		"status": "error",
		"error": gin.H{
			"code":    code,
			"message": message,
		},
	})
	c.Abort()
}

// JSONSuccess is a unified success wrapper.
// Pass any payload.
func JSONSuccess(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, gin.H{
		"status": "success",
		"data":   payload,
	})
}

// JSONValidationError helps return field-specific validation errors.
func JSONValidationError(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"error": gin.H{
			"code":   "VALIDATION_ERROR",
			"fields": errors,
		},
	})
	c.Abort()
}
