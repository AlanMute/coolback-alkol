package service

import "github.com/KrizzMU/coolback-alkol/internal/repository"

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Add(name string, description string, folderName string) error {

	return s.repo.Add(name, description, folderName)
}
