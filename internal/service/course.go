package service

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

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
	lessonsToDelete, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	for _, lessonToDelete := range lessonsToDelete {
		fileName := strconv.FormatUint(uint64(lessonToDelete), 10) + ext

		pathToFile := filepath.Join("./lessons", fileName)
		fmt.Println(pathToFile)

		if err := os.Remove(pathToFile); err != nil {
			return err
		}
	}

	return nil
}

func (s *CourseService) GetByName(name string) ([]core.Course, error) {

	return s.repo.GetByName(name)
}

func (s *CourseService) GetAll() ([]core.Course, error) {

	return s.repo.GetAll()
}

func (s *CourseService) Get(id int) (core.CourseContent, error) {
	content, err := s.repo.Get(id)

	if err != nil {
		return content, err
	}

	return content, nil
}

func (s *CourseService) Put(id int, name string, desc string) error {
	if err := s.repo.Put(id, name, desc); err != nil {
		return err
	}

	return nil
}
