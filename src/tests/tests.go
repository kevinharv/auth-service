package tests

import (
	"fmt"
	"kevinharv/auth-service/src/db"
	"kevinharv/auth-service/src/utils"
)

// Insert MockSAML IdP
func InsertIDP() {
	db, err := db.Connect()
	utils.HandleErr(err, "DB connection failed during testing.")
	defer db.Close()

	// Check if IdP exists

	// If not, insert
	res, err := db.Exec("INSERT INTO saml_idps (display_name, domain, metadata_url) VALUES ($1, $2, $3)", "MockSAML", "example.com", "https://mocksaml.com/api/saml/metadata")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", res)
}

// Insert Authentication Method
func InsertAuthMeth() {

}

// Insert Users
func InsertUsers() {

}
