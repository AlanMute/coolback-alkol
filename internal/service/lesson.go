package service

import (
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

func (s *LessonService) Add(name string, description string, id int) error {
	var path string // ex. "/couse_1/module_1"

	fileName, err := pkg.UniqueFile(ext, "/courses"+path)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, id, fileName)
}

func (s *LessonService) Get(name string) error {
	return s.Get(name)
}
