package repository

import "github.com/jinzhu/gorm"

type Repository interface {
	CloseConnection()
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
