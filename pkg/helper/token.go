package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

// TokenInfo holds extracted token data
type TokenInfo struct {
	UserID     string `json:"user_id"`
	ClientType string `json:"client_type"`
}

// GenerateJWT creates a signed JWT token with provided claims, expiration time, and secret key.
func GenerateJWT(claimsMap map[string]interface{}, expireDuration time.Duration, secretKey string) (string, error) {
	claims := jwt.MapClaims{}

	// Add provided claims to token
	for k, v := range claimsMap {
		claims[k] = v
	}

	now := time.Now()
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(expireDuration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseClaims parses the JWT token string using secret key and extracts TokenInfo
func ParseClaims(tokenString string, secretKey string) (TokenInfo, error) {
	claims, err := extractClaims(tokenString, secretKey)
	if err != nil {
		return TokenInfo{}, err
	}

	userID := cast.ToString(claims["user_id"])
	if userID == "" {
		return TokenInfo{}, errors.New("cannot parse 'user_id' field")
	}

	clientType := cast.ToString(claims["client_type"])

	return TokenInfo{
		UserID:     userID,
		ClientType: clientType,
	}, nil
}

// extractClaims extracts and validates claims from the JWT token string
func extractClaims(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check token signing method for security
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// ExtractToken extracts token part from "Bearer <token>" authorization header string
func ExtractToken(bearer string) (string, error) {
	parts := strings.SplitN(bearer, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
