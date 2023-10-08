package repository

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type Course interface {
	Add(name string, description string, folderName string) error
	Delete(id uint) (string, error)
	GetByName(name string) ([]core.Course, error)
	GetAll() ([]core.Course, error)
	Get(id int) (core.СourseСontent, error)
	Put(id int, name string, desc string) error
}

type Module interface {
	Add(name string, description string, courseName string, folderName string) error
	Put(id int, name string, desc string, orderid uint) error
	Delete(id uint) (string, error)
	Get(id int) (core.ModLes, error)
}

type Lesson interface {
	Add(name string, description string, fileName string, courseName string, moduleName string) error
	Delete(id uint) (string, error)
	Get(moduleid int, orderid int) (core.Lesson, error)
	Put(id int, name string, desc string, orderID uint) error
}

type Session interface {
	Add(session core.Sessions) error
	CheckRefresh(token string) error
}

type Repository struct {
	Course
	Module
	Lesson
	Session
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Course:  NewCoursePostgres(db),
		Module:  NewModulePostgres(db),
		Lesson:  NewLessonPostgres(db),
		Session: NewSessionPostgres(db),
	}
}
