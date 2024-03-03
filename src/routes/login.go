package routes

import (
	"net/http"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// Handle login flow
func HandleLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		dest := c.Request.URL.Query().Get("dest")
		if dest == "" {
			dest =  os.Getenv("DEFAULT_REDIR_DEST")
		}
		
		c.SetCookie("auth_dest", dest, 60, "/", "localhost", true, true)
		
		upn := c.Request.URL.Query().Get("upn")
		if upn == "" {
			c.Status(http.StatusBadRequest)
			return
		}

        fmt.Printf("User Auth Method: %s", upn)

        // Lookup auth method in DB
        // If SAML
            // Load SP
            // Handle Start Auth Flow
            // IN ACS - create and issue token

        // If !SAML
            // Handle OAuth for MSFT or GOOG
            // For now, just exit with internal server error
        

        c.Status(http.StatusOK)
		//tk, err := utils.GenerateJWT("test@test.com")
		//if err != nil {
		//	c.Status(http.StatusInternalServerError)
		//	return
		//}

		// Lookup auth method by UPN
		// Direct to auth method
		//c.JSON(http.StatusOK, gin.H{"token": tk})
	}
}
