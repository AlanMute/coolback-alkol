package repository

import "github.com/jinzhu/gorm"

type Repository interface {
	AddCourse(name string, description string, folderName string) error
	AddModule(name string, description string, id int, folderName string) error
	AddLesson(name string, description string, id int, fileName string) error
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
