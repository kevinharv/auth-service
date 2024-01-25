package main

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	
	"github.com/crewjam/saml/samlsp"
	"github.com/gin-gonic/gin"
)

func GinHandlerFromHTTPHandler(httpHandler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a fake ResponseWriter to pass to the HTTP handler
		fakeWriter := &fakeResponseWriter{c.Writer, http.StatusOK}
		// Call the HTTP handler with the fake ResponseWriter
		httpHandler.ServeHTTP(fakeWriter, c.Request)
		// Set the status code from the fake ResponseWriter to the Gin context
		c.Status(fakeWriter.status)
	}
}

// A fakeResponseWriter is a simple wrapper around gin.ResponseWriter
// that allows capturing the HTTP status code.
type fakeResponseWriter struct {
	gin.ResponseWriter
	status int
}

// WriteHeader captures the HTTP status code.
func (w *fakeResponseWriter) WriteHeader(code int) {
	w.status = code
}

func main() {
	keyPair, err := tls.LoadX509KeyPair("cert/myservice.cert", "cert/myservice.key")
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	// idpMetadataURL, err := url.Parse("https://samltest.id/saml/idp")
	idpMetadataURL, err := url.Parse("https://mocksaml.com/api/saml/metadata")
	if err != nil {
		panic(err) // TODO handle error
	}
	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
		*idpMetadataURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	rootURL, err := url.Parse("http://localhost:8000")
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	router := gin.Default()

	fmt.Printf(samlSP.ServiceProvider.EntityID)

	httpHandler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from HTTP handler!")
	})

	// Use the GinHandlerFromHTTPHandler middleware to convert the HTTP handler to a Gin handler
	ginHandler := GinHandlerFromHTTPHandler(samlSP.RequireAccount(httpHandler))

	// Define a Gin route that uses the converted Gin handler
	router.GET("/hello", ginHandler)

	// router.Use(func(c *gin.Context) {
	// 	samlSP.ServeHTTP(c.Writer, c.Request)
	// 	fmt.Printf("Testing\n")
	// 	c.Next()
	// })

	router.POST("/saml/acs", func(c *gin.Context) {
		// samlSP.ServeACS(c.Writer, c.Request)
		samlSP.ServeACS(c.Writer, c.Request)
		c.Redirect(http.StatusFound, samlSP.ServiceProvider.DefaultRedirectURI)
		c.Abort()
	})

	router.GET("/saml", func(c *gin.Context) {
		samlSP.HandleStartAuthFlow(c.Writer, c.Request)
		c.Abort()
	})

	// Protected route
	router.GET("/test", func(c *gin.Context) {
		fmt.Printf("In route\n")
		c.JSON(200, gin.H{"message": "Protected resource"})
	})


	// Run the application on port 8080
	router.Run(":8000")
}

// Gin handler
// HTTP handler