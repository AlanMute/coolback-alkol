package service

import (
	"mime/multipart"
	"os"
	"path/filepath"

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

	return s.repo.Add(name, description, dbfileName, courseName, moduleName)
}

func (s *LessonService) Delete(name string, courseName string, moduleName string) error {
	coursePath, err := pkg.GetPath(courseName, "./courses")
	if err != nil {
		return err
	}

	modulePath, err := pkg.GetPath(moduleName, coursePath)
	if err != nil {
		return err
	}

	filePath, err := pkg.GetPathToFile(name, ext, modulePath)
	if err != nil {
		return err
	}

	if err := os.Remove(filePath); os.IsNotExist(err) {
		return err
	}

	return s.repo.Delete(name, courseName, moduleName)
}

func (s *LessonService) Get(course string, module string, lesson string) ([]string, error) {
	path := filepath.Join("./courses", course, module, lesson+".md")

	strFile, err := pkg.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return strFile, nil
}
