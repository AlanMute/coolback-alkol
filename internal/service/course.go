package service

import (
	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Add(name string, description string) error {
	if err := s.repo.Add(name, description); err != nil {
		return err
	}

	return nil
}

func (s *CourseService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *CourseService) GetByName(name string) ([]core.Course, error) {

	return s.repo.GetByName(name)
}

func (s *CourseService) GetAll() ([]core.Course, error) {

	return s.repo.GetAll()
}

func (s *CourseService) Get(name string) (core.CourseContent, error) {

	path := "courses\\" + name

	content, err := s.repo.Get(path)

	if err != nil {
		return content, err
	}

	return content, nil
}
