/*
The SAML component of the middleware package contains functions
to perform authentication verification and validation. These are
implemented as Gin middleware handler functions.
*/

package middleware

import (
	"net/http"
	"github.com/crewjam/saml/samlsp"
	"github.com/gin-gonic/gin"
)

// SAML Authentication Middleware
func SAMLMiddleware(sp *samlsp.Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := sp.Session.GetSession(c.Request)
		if err != nil {
			c.Redirect(http.StatusFound, "/saml/login")
		}
		c.Next()
	}
}