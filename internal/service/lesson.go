package service

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg"
)

const (
	ext string = ".md"
)

type LessonService struct {
	repo repository.Lesson
}

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

	if err := s.repo.Add(name, description, dbfileName, courseName, moduleName); err != nil {
		if rmErr := os.Remove(dbfileName); rmErr != nil {
			return rmErr
		}

		return err
	}

	return nil
}

func (s *LessonService) Delete(id uint) error {
	filePath, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	if err := os.Remove(filePath); !os.IsNotExist(err) {
		fmt.Printf("err = %e", err)
		return err
	}

	return nil
}

func (s *LessonService) Get(name string) error {
	return s.repo.Get(name)
}
