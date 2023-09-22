package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
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

type Service struct {
	Course
	Module
	Lesson
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Course: NewCourseService(repo.Course),
		Module: NewModuleService(repo.Module),
		Lesson: NewLessonService(repo.Lesson),
	}
}
