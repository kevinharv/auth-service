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
		authorizationHeader := c.Request.Header.Get("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not supplied in request."})
			c.Abort()
			return
		}

		headerStringArr := strings.Split(authorizationHeader, " ")
		if len(headerStringArr) < 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Malformed authorization header"})
			c.Abort()
			return
		}

		reqJWT := headerStringArr[1]
		fmt.Printf("DEBUG: Request JWT: %s\n", reqJWT)

		token, err := utils.ParseToken(reqJWT)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return

		}

		claims, err := utils.ParseClaims(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Claims"})
			c.Abort()
			return
		}
		
		/*
			TODO - Finish Implementation
			- get auth method by user claim
			- direct to auth method
		*/

		fmt.Printf("DEBUG: userPrincipalName: %s\n", claims.UserPrincipalName)
	}
}
