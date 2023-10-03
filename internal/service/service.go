package service

import (
	"mime/multipart"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type Course interface {
	Add(name string, description string) error
	Delete(id uint) error
	GetByName(name string) ([]core.Course, error)
	GetAll() ([]core.Course, error)
	Get(name string) (core.CourseContent, error)
}

type Module interface {
	Add(name string, description string, orderID uint, courseName string) error
	Delete(id uint) error
	Get(moduleName string, courseName string) (core.ModLes, error)
}

type Lesson interface {
	Add(file multipart.File, fileName string, name string, description string, orderID uint, moduleName string, courseName string) error
	Delete(id uint) error
	Get(course string, module string, lesson string) (core.LesMd, error)
	SendTrialLesson(email string) error
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
