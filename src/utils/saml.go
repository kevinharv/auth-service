/*
Module for SAML Service Provider configuration logic.
*/

package utils

import (
	"context"
	"crypto/tls"
	"crypto/rsa"
	"crypto/x509"
	"net/http"
	"net/url"
	
	"github.com/crewjam/saml/samlsp"
)

const (
	serviceProviderKeyPath  = "cert/myservice.key"
	serviceProviderCertPath = "cert/myservice.cert"
	idpMetadataURL          = "https://mocksaml.com/api/saml/metadata"
	rootURL                 = "http://localhost:8080"
)

func InitSAMLSP() samlsp.Middleware {
	keyPair, err := tls.LoadX509KeyPair(serviceProviderCertPath, serviceProviderKeyPath)
	HandleErr(err, "Failed to load SAML SP X.509 Key Pair")

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	HandleErr(err, "Failed to parse SAML SP Certificate")

	idpMetadataURL, err := url.Parse(idpMetadataURL)
	HandleErr(err, "Failed to parse SAML IdP Metadata URL")

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	HandleErr(err, "Failed to load SAML IdP Metadata")

	rootURL, err := url.Parse(rootURL)
	HandleErr(err, "Failed to parse SAML SP root URL")

	samlSP, err := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	HandleErr(err, "Failed to setup SAML SP")
	return *samlSP
}