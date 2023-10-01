package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/KrizzMU/coolback-alkol/internal/core"
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

	dbfileName, err := pkg.GenerateUniqueFile(fileName, name, path, ext)
	if err != nil {
		return err
	}

	if err := s.repo.Add(name, description, dbfileName, courseName, moduleName); err != nil {
		return err
	}

	if err := pkg.CreateFile(file, dbfileName); err != nil {
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

func (s *LessonService) Get(moduleid int, orderid int) (core.LesMd, error) {
	var lesmd core.LesMd

	lesson, err := s.repo.Get(moduleid, orderid)

	if err != nil {
		return lesmd, err
	}

	path := filepath.Join("lessons", fmt.Sprint(lesson.ID)+".md")

	file, err := pkg.ReadFile(path)

	if err != nil {
		return lesmd, err
	}

	lesmd = core.LesMd{
		Lesson: lesson,
		Mdfile: file,
	}

	return lesmd, nil
}

func (s *LessonService) Put(id int, name string, desc string, orderID uint, content []string) error {
	if len(content) > 0 {
		path := filepath.Join("lessons", fmt.Sprint(id)+".md")
		if err := pkg.UpdateFile(path, content); err != nil {
			return err
		}
	}

	err := s.repo.Put(id, name, desc, orderID)

	if err != nil {
		return err
	}

	return nil
}
