package service

import (
	"fmt"
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

func (s *LessonService) Add(name string, description string, orderID uint, moduleID uint, content []string) error {

	if name == "" {
		name = "New lesson " + fmt.Sprint(orderID)
	}

	lessonID, err := s.repo.Add(name, description, orderID, moduleID)
	if err != nil {
		return err
	}

	filePath := filepath.Join("./lessons", fmt.Sprint(lessonID)+ext)

	if err := pkg.CreateFile(filePath, content); err != nil {
		_, errDel := s.repo.Delete(lessonID)
		if errDel != nil {
			return errDel
		}
		return err
	}

	return nil
}

func (s *LessonService) Delete(id uint) error {
	fileID, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	fileName := fileID + ext

	filePath := filepath.Join("./lessons", fileName)

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

	path := filepath.Join("lessons", fmt.Sprint(lesson.ID)+ext)

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
		path := filepath.Join("lessons", fmt.Sprint(id)+ext)
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
