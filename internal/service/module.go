package service

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/internal/repository"
)

type ModuleService struct {
	repo repository.Module
}

func NewModuleService(repo repository.Module) *ModuleService {
	return &ModuleService{repo: repo}
}

// Пофиксить возможность добавления разных модулей одного курса с одинаковым orderID

func (s *ModuleService) Add(name string, description string, orderID uint, courseName string) error {
	if err := s.repo.Add(name, description, orderID, courseName); err != nil {
		return err
	}

	return nil
}

func (s *ModuleService) Delete(id uint) error {
	lessonsToDelete, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	for _, lessonToDelete := range lessonsToDelete {
		FileName := strconv.FormatUint(uint64(lessonToDelete), 10) + ext

		pathToLesson := filepath.Join("./lessons", FileName)

		if err := os.Remove(pathToLesson); err != nil {
			return err
		}
	}

	return nil
}

func (s *ModuleService) Get(moduleName string, courseName string) (core.ModLes, error) {
	path := "courses\\" + courseName + "\\" + moduleName
	modles, err := s.repo.Get(path)
	if err != nil {
		return modles, err
	}

	return modles, nil
}
