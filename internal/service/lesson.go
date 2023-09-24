package service

import (
	"mime/multipart"

	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg"
)

type LessonService struct {
	repo repository.Lesson
}

var ext string = ".md"

func NewLessonService(repo repository.Lesson) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) Add(file multipart.File, fileName string, name string, description string, moduleName string, courseName string) error {
	coursePath, err := pkg.GetPath(courseName, "./courses")
	if err != nil {
		return err
	}

	path, err := pkg.GetPath(moduleName, coursePath)
	if err != nil {
		return err
	}

	dbfileName, err := pkg.CreateUniqueFile(file, fileName, name, path, ext)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, dbfileName, courseName, moduleName)
}

func (s *LessonService) Get(name string) error {
	return s.repo.Get(name)
}
