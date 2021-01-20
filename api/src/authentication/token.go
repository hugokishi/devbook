package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateToken - GenerateToken for user
func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken - Validate JWT token
func ValidateToken(r *http.Request) error {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, verificationKey)
	if err != nil {
		return err
	}

	if _, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		return nil
	}

	return errors.New("Invalid Token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func verificationKey(token *jwt.Token) (interface{}, error) {
	if _, err := token.Method.(*jwt.SigningMethodHMAC); !err {
		return nil, fmt.Errorf("Signature method invalid! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// ExtractUserID - Return UserID saved on token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenStr := extractToken(r)
	token, err := jwt.Parse(tokenStr, verificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")
}
