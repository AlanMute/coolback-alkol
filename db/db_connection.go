package db

import (
	"database/sql"

	"github.com/KrizzMU/coolback-alkol/config"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", config.GetConnectionString())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
