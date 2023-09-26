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
	dbFolderName, err := pkg.CreateUniqueFolder(name, "./courses")
	if err != nil {
		return err
	}

	if err := s.repo.Add(name, description, dbFolderName); err != nil {
		if rmErr := os.RemoveAll(dbFolderName); rmErr != nil {
			return rmErr
		}

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
