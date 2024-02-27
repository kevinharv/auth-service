package main

import (
	"net/http"
	"time"

	"kevinharv/auth-service/src/db"
	"kevinharv/auth-service/src/middleware"
	"kevinharv/auth-service/src/utils"

	"github.com/gin-gonic/gin"
)



func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	defer db.Close()


	// Setup SAML Service Provider
	sp := utils.InitSAMLSP()

	// Setup Gin Router
	r := gin.Default()

	// SAML Routes
	r.GET("/saml/login", func(c *gin.Context) {
		sp.HandleStartAuthFlow(c.Writer, c.Request)
	})
	r.POST("/saml/acs", func(c *gin.Context) {
		sp.ServeACS(c.Writer, c.Request)
		c.Redirect(301, "/hello")
	})
	r.GET("/saml/metadata", func(c *gin.Context) {
		spMetadata := sp.ServiceProvider.Metadata()
		c.XML(200, spMetadata)
	})
	r.GET("/saml/logout", func(c *gin.Context) {
		sp.Session.DeleteSession(c.Writer, c.Request)
		c.JSON(200, gin.H{"status": "ok"})
	})

	// SAML Protected Routes
	authorized := r.Group("/")
	authorized.Use(middleware.SAMLMiddleware(&sp))
	{
		authorized.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"hello": "world"})
		})
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}