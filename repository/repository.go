package repository

import "database/sql"

type Repository interface {
	CloseConnection()
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
