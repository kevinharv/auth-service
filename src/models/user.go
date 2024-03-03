/*
This module contains database models for user management.
*/

package user

import (
	"kevinharv/auth-service/src/db"
	"kevinharv/auth-service/src/utils"
)

type User struct {
	userID            string
	userPrincipalName string
	authMethod        string
	firstName         string
	lastName          string
	middleInit        string
	displayName       string
}

// Create User
func CreateUser(u *User) {
	db, err := db.Connect()
	utils.HandleErr(err, "Failed to connect to DB.")
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (userPrincipalName, auth_method, first_name, last_name, middle_init, display_name) VALUES ($1, $2, $3, $4, $5, $6)", 
					u.userPrincipalName, u.authMethod, u.firstName, u.lastName, u.middleInit, u.displayName)
	utils.HandleErr(err, "Failed to insert user into DB.")
}

// Delete User by ID
func DeleteUser(userID string) {
	db, err := db.Connect()
	utils.HandleErr(err, "Failed to connect to DB.")
	defer db.Close()

	_, err = db.Exec("DELETE FROM users WHERE user_id = '$1'", userID)
	utils.HandleErr(err, "Failed to delete user from DB.")
}

// Get User by ID
func GetUser(userID string) User {
	db, err := db.Connect()
	utils.HandleErr(err, "Failed to connect to DB.")
	defer db.Close()

	rows, err := db.Query("SELECT (user_id, userPrincipalName, auth_method, first_name, last_name, middle_init, display_name) FROM users WHERE user_id = '$1' LIMIT 1", userID)
	utils.HandleErr(err, "Failed to get user from DB.")
	defer rows.Close()

	var u User

	for rows.Next() {
		rows.Scan(&u.userID, &u.userPrincipalName, &u.authMethod, &u.firstName, &u.lastName, &u.middleInit, &u.displayName)
	}

	return u
}

// Get User by UPN
func GetUserByUPN(upn string) User {
	db, err := db.Connect()
	utils.HandleErr(err, "Failed to connect to DB.")
	defer db.Close()

	rows, err := db.Query("SELECT (user_id, userPrincipalName, auth_method, first_name, last_name, middle_init, display_name) FROM users WHERE userPrincipalName = '$1' LIMIT 1", upn)
	utils.HandleErr(err, "Failed to get user from DB.")
	defer rows.Close()

	var u User

	for rows.Next() {
		rows.Scan(&u.userID, &u.userPrincipalName, &u.authMethod, &u.firstName, &u.lastName, &u.middleInit, &u.displayName)
	}

	return u
}