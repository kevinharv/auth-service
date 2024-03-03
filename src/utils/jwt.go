package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

// JWT claim struct - define fields to be carried in token
// Use additional fields for authorization, user info, etc.
type Claims struct {
	UserPrincipalName string `json:"userPrincipalName"`
	jwt.RegisteredClaims
}

// Generate a JWT with only a UPN - TODO REQUIRE MORE INFO
func GenerateJWT(upn string) (string, error) {
	claims := Claims{
		upn,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "Authentication Microservice",
			Subject:   upn,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Parse token from string - will throw error if expired or tampered with
func ParseToken(encToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encToken, &Claims{}, func(token *jwt.Token) (interface{}, error) { return secretKey, nil})
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

// Retrieve the claims from the token
func ParseClaims(token *jwt.Token) (*Claims, error) {
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("failed to extract JWT claims")
	}
	
	return claims, nil
}