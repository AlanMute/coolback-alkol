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

func (s *LessonService) Add(file multipart.File, name string, description string, id int) error {
	var path string // ex. "/couse_1/module_1"

	fileName, err := pkg.CreateUniqueFile(file, name, path, ext)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, id, fileName)
}

func (s *LessonService) Get(name string) error {
	return s.repo.Get(name)
}
