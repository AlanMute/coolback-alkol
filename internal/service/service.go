package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
)

const (
	lessonExt = ".md"
	imageExt  = ".jpg"
)

type Course interface {
	Add(name string, description string) error
	Delete(id uint) error
	GetByName(name string) ([]core.Course, error)
	GetAll() ([]core.Course, error)
	Get(id int) (core.CourseContent, error)
	Put(id int, name string, desc string) error
	GetImage(id uint) (string, error)
	DeleteImage(id uint) error
}

type Module interface {
	Add(name string, description string, orderID uint, courseID uint) error
	Delete(id uint) error
	Get(id int) (core.ModLes, error)
	Put(id int, name string, desc string, orderid uint) error
	GetImage(id uint) (string, error)
	DeleteImage(id uint) error
}

type Lesson interface {
	Add(name string, description string, orderID uint, moduleID uint, content []string) error
	Delete(id uint) error
	Get(moduleid int, orderid int) (core.LesMd, error)
	Put(id int, name string, desc string, orderID uint, content []string) error
	SendTrialLesson(email string) error
}

type Admin interface {
	SignIn(login string, pass string) (core.Tokens, error)
	Refresh(refreshToken string) (string, error)
}

type Service struct {
	Course
	Module
	Lesson
	Admin
}

func NewService(repo *repository.Repository, t auth.TokenManager) *Service {
	return &Service{
		Course: NewCourseService(repo.Course),
		Module: NewModuleService(repo.Module),
		Lesson: NewLessonService(repo.Lesson),
		Admin:  NewAdminService(t, repo.Session),
	}
}
