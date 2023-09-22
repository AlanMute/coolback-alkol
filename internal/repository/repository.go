package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Course interface {
	Add(name string, description string, folderName string) error
	GetByName(name string) ([]core.Course, error)
	GetAll() ([]core.Course, error)
}

type Module interface {
	Add(name string, description string, id int, folderName string) error
}

type Lesson interface {
	Add(name string, description string, id int, fileName string) error
	Get(name string) error
}

type Repository struct {
	Course
	Module
	Lesson
	//db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Course: NewCoursePostgres(db),
		Module: NewModulePostgres(db),
		Lesson: NewLessonPostgres(db),
	}
}
