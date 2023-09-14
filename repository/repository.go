package repository

import "database/sql"

type Repository interface {
	AddCourse(name string, description string, folderName string) error
	AddModule(name string, description string, id int, folderName string) error
	AddLesson(name string, description string, id int, fileName string) error
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
