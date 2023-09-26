package service

import (
	"os"

	"github.com/KrizzMU/coolback-alkol/internal/repository"
	"github.com/KrizzMU/coolback-alkol/pkg"
)

type ModuleService struct {
	repo repository.Module
}

func NewModuleService(repo repository.Module) *ModuleService {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) Add(name string, description string, courseName string) error {
	path, err := pkg.GetPath(courseName, "./courses")
	if err != nil {
		return err
	}

	dbFolderName, err := pkg.CreateUniqueFolder(name, path)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, courseName, dbFolderName)
}

func (s *ModuleService) Delete(id uint) error {
	dirPath, err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(dirPath); !os.IsNotExist(err) {
		return err
	}

	return nil
}
