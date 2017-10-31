package main

import (
	"database/sql"			// package SQL
	_ "github.com/lib/pq"	// driver Postgres
	"fmt"
)

const (
	DB_USER     = "poe"
	DB_PASSWORD = "poepass"
	DB_NAME     = "poe"
	DB_HOST     = "127.0.0.1"
)

/*
* dbname - The name of the database to connect to
* user - The user to sign in as
* password - The user's password
* host - The host to connect to. Values that start with / are for unix domain sockets. (default is localhost)
* port - The port to bind to. (default is 5432)
* sslmode - Whether or not to use SSL (default is require, this is not the default for libpq)
* fallback_application_name - An application_name to fall back to if one isn't provided.
* connect_timeout - Maximum wait for connection, in seconds. Zero or not specified means wait indefinitely.
* sslcert - Cert file location. The file must contain PEM encoded data.
* sslkey - Key file location. The file must contain PEM encoded data.
* sslrootcert - The location of the root certificate file. The file must contain PEM encoded data.
 */

func connectDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD,DB_HOST ,DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	return db,err
}
