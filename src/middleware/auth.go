package middleware

import (
	"fmt"
	"kevinharv/auth-service/src/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization Header
		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not supplied in request."})
			c.Abort()
			return
		}

		// Check for OAuth mechanism spec
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed Authorization header"})
			c.Abort()
			return
		}

		// Extract the token from the header
		reqToken := strings.TrimPrefix(authorizationHeader, "Bearer ")

		// Parse the token, check for validity
		token, err := utils.ParseToken(reqToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return

		}

		// Get the claims - will can carry user info and authorization info
		claims, err := utils.ParseClaims(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Claims"})
			c.Abort()
			return
		}

		fmt.Printf("DEBUG: userPrincipalName: %s\n", claims.UserPrincipalName)
	}
}
