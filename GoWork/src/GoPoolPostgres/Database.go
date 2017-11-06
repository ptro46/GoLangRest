package main

import (
	"database/sql" // package SQL
	"fmt"
	_ "github.com/lib/pq" // driver Postgres
)

const (
	DB_USER     = "poe"
	DB_PASSWORD = "poepass"
	DB_NAME     = "poe"
	DB_HOST     = "127.0.0.1"
)

type Database struct {
	db *sql.DB
}

func connectDB() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(4)
	return db, err
}
