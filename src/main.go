package main

import (
	"net/http"
	"time"

	"kevinharv/auth-service/src/middleware"
	"kevinharv/auth-service/src/routes"
	"kevinharv/auth-service/src/tests"
	"kevinharv/auth-service/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Insert test data into the database
	tests.InsertIDP()
	tests.InsertAuthMeth()
	tests.InsertUsers()

	// Setup Gin Router
	r := gin.Default()

	// Setup SAML Service Provider
	sp := utils.InitSAMLSP()

	// SAML Routes
	r.GET("/saml/login", func(c *gin.Context) {
		sp.HandleStartAuthFlow(c.Writer, c.Request)
	})
	r.POST("/saml/acs", func(c *gin.Context) {
		sp.ServeACS(c.Writer, c.Request)

		dest, err := c.Cookie("auth_dest")
		if err != nil {
			dest = "localhost:8000"
		}
		
		c.SetCookie("auth_dest", "", 0, "/", "localhost", true, true)
		c.Redirect(301, dest)
	})
	r.GET("/saml/metadata", func(c *gin.Context) {
		spMetadata := sp.ServiceProvider.Metadata()
		c.XML(200, spMetadata)
	})
	r.GET("/saml/logout", func(c *gin.Context) {
		sp.Session.DeleteSession(c.Writer, c.Request)
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/login", routes.HandleLogin())

	r.GET("/token", func(c *gin.Context) {
		jwt, err := utils.GenerateJWT("test@test.com")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT"})
		}
		c.JSON(http.StatusOK, gin.H{"token": jwt})
	})

	jwtd := r.Group("/jwt")
	jwtd.Use(middleware.JWTMiddleware())
	{
		jwtd.GET("/protected", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"route": "yay"})
		})
	}

	// SAML Protected Routes
	authorized := r.Group("/")
	authorized.Use(middleware.SAMLMiddleware(&sp))
	{
		authorized.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"hello": "world"})
		})
	}

	s := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
