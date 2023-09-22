package service

import (
	"os"
	"strings"

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
	folderName, err := pkg.UniqueFolder("course", "./courses")
	if err != nil {
		return err
	}

	if err := os.Mkdir(folderName, os.ModePerm); err != nil {
		return err
	}

	splitFileName := strings.Split(folderName, "\\")
	dbFolderName := splitFileName[len(splitFileName)-1]

	return s.repo.Add(name, description, dbFolderName)
}

func (s *CourseService) GetByName(name string) ([]core.Course, error) {

	return s.repo.GetByName(name)
}

func (s *CourseService) GetAll() ([]core.Course, error) {

	return s.repo.GetAll()
}
