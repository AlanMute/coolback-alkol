package service

import "github.com/KrizzMU/coolback-alkol/internal/repository"

type LessonService struct {
	repo repository.Lesson
}

func NewLessonService(repo repository.Lesson) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) Add(name string, description string, id int, fileName string) error {
	return s.repo.Add(name, description, id, fileName)
}

func (s *LessonService) Get(name string) error {
	return s.Get(name)
}
