package service

import (
	"os"
	"path/filepath"

	"github.com/KrizzMU/coolback-alkol/internal/core"
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

	dbFolderName, err := pkg.GenerateUniqueFolder(name, path)
	if err != nil {
		return err
	}

	if err := s.repo.Add(name, description, courseName, dbFolderName); err != nil {
		return err
	}

	if err := pkg.CreateFolder(dbFolderName); err != nil {
		return err
	}

	return nil
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

func (s *ModuleService) Get(moduleName string, courseName string) (core.ModLes, error) {
	path := filepath.Join(courseName, moduleName)
	modles, err := s.repo.Get(path)
	if err != nil {
		return modles, err
	}

	return modles, nil
}
