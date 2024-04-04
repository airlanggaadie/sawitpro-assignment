package main

import (
	"database/sql"
	"fmt"
	"os"
)

func initPostgresql() (*sql.DB, error) {
	dbDsn := os.Getenv("DATABASE_URL")
	if dbDsn == "" {
		return nil, fmt.Errorf("please configure DATABASE_URL environment variable")
	}

	db, err := sql.Open("postgres", dbDsn)
	if err != nil {
		return nil, fmt.Errorf("[main][initPostgresql] open database error: %v", err)
	}

	return db, nil
}
