/*
This module contains functions related to database connectivity. It includes
configuration items for the database connection, and helper functions for 
interaction.
*/

package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	hostname = "localhost"
	username = "postgres"
	password = "postgres"
	database = "attn"
)

func Connect() (*sql.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, hostname, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
