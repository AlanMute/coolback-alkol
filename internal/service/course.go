package service

import (
	"os"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg"
)

type CourseService struct {
	repo repository.Course
}

func NewCourseService(repo repository.Course) *CourseService {
	return &CourseService{repo: repo}
}

func (s *CourseService) Add(name string, description string) error {
	dbFolderName, err := pkg.GenerateUniqueFolder(name, "./courses")
	if err != nil {
		return err
	}

	if err := s.repo.Add(name, description, dbFolderName); err != nil {
		return err
	}

	if err := pkg.CreateFolder(dbFolderName); err != nil {
		return err
	}

	return nil
}

func (s *CourseService) Delete(id uint) error {
	dirPath, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(dirPath); !os.IsNotExist(err) {
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

func (s *CourseService) Get(name string) (core.СourseСontent, error) {

	path := "courses\\" + name

	content, err := s.repo.Get(path)

	if err != nil {
		return content, err
	}

	return content, nil
}
