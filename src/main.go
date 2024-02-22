package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/crewjam/saml/samlsp"
	"github.com/gin-gonic/gin"
)

const (
	serviceProviderKeyPath  = "cert/myservice.key"
	serviceProviderCertPath = "cert/myservice.cert"
	idpMetadataURL          = "https://mocksaml.com/api/saml/metadata"
	rootURL                 = "http://localhost:8080"
)

func initSAMLSP() samlsp.Middleware {
	keyPair, err := tls.LoadX509KeyPair(serviceProviderCertPath, serviceProviderKeyPath)
	handleErr(err, "Failed to load SAML SP X.509 Key Pair")

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	handleErr(err, "Failed to parse SAML SP Certificate")

	idpMetadataURL, err := url.Parse(idpMetadataURL)
	handleErr(err, "Failed to parse SAML IdP Metadata URL")

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	handleErr(err, "Failed to load SAML IdP Metadata")

	rootURL, err := url.Parse(rootURL)
	handleErr(err, "Failed to parse SAML SP root URL")

	samlSP, err := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	handleErr(err, "Failed to setup SAML SP")
	return *samlSP
}

func samlMiddleware(sp *samlsp.Middleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := sp.Session.GetSession(c.Request)
		if err != nil {
			fmt.Printf("INFO: SAML Session Missing\n")
			c.Redirect(http.StatusFound, "/saml/login")
		}
		c.Next()
	}
}

func main() {
	// Setup SAML Service Provider
	sp := initSAMLSP()

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

	// SAML Protected Routes
	authorized := r.Group("/")
	authorized.Use(samlMiddleware(&sp)) 
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

func handleErr(e error, msg string) {
	if e != nil {
		fmt.Printf("ERROR: %s\n", msg)
		panic(e)
	}
}
