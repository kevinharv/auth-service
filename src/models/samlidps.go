package models

import (
	"net/url"
)

func GetIdPMetadataURL(idp_id string) *url.URL {
    // Query DB for url
   // Return URL
   mdurl, err := url.Parse("https://localhost:8000")
   if err != nil {
       panic(err)
   }

   return mdurl
}
