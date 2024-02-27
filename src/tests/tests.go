package tests

import (
	"kevinharv/auth-service/src/db"
	"kevinharv/auth-service/src/utils"
)

// Insert MockSAML IdP
func InsertIDP() {
	db, err := db.Connect()
	utils.HandleErr(err, "DB connection failed during testing.")
	defer db.Close()

	// Check if IdP exists
	rows, err := db.Query("SELECT * FROM saml_idps")
	utils.HandleErr(err, "Failed to query DB for IdPs")
	defer rows.Close()

	var rowCount int
	for rows.Next() {
		rowCount++
	}

	// If not, insert IdP
	if rowCount == 0 {
		_, err := db.Exec("INSERT INTO saml_idps (display_name, domain, metadata_url) VALUES ($1, $2, $3)", "MockSAML", "example.com", "https://mocksaml.com/api/saml/metadata")
		if err != nil {
			panic(err)
		}
	}
}

// Insert Authentication Method
func InsertAuthMeth() {
	db, err := db.Connect()
	utils.HandleErr(err, "DB connection failed during testing.")
	defer db.Close()

	// Check if IdP exists
	rows, err := db.Query("SELECT * FROM auth_methods")
	utils.HandleErr(err, "Failed to query DB for authentication methods")
	defer rows.Close()

	var rowCount int
	for rows.Next() {
		rowCount++
	}

	// If not, insert IdP
	if rowCount == 0 {
		_, err := db.Exec("INSERT INTO auth_methods (auth_name) VALUES ($1)", "SAML 2.0")
		if err != nil {
			panic(err)
		}
	}
}

// Insert Users
func InsertUsers() {
	db, err := db.Connect()
	utils.HandleErr(err, "DB connection failed during testing.")
	defer db.Close()

	usersRows, err := db.Query("SELECT * FROM users")
	utils.HandleErr(err, "DB query failed during user testing.")

	var rowCount int
	for usersRows.Next() {
		rowCount++
	}

	if rowCount != 0 {
		return
	}

	var authMethod string // Get from DB
	upn := "test@test.com"
	fname := "Testy"
	lname := "McTestFace"
	mi := "T"
	displayName := "McTestFace, Testy"


	// Check if IdP exists
	rows, err := db.Query("SELECT auth_id FROM auth_methods")
	utils.HandleErr(err, "Failed to query DB for authentication methods")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&authMethod)
	}

	_, err = db.Exec("INSERT INTO users (userPrincipalName, auth_method, first_name, last_name, middle_init, display_name) VALUES ($1, $2, $3, $4, $5, $6)", upn, authMethod, fname, lname, mi, displayName)
	utils.HandleErr(err, "Failed to insert users during testing.")
}
