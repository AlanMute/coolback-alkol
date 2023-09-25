package service

import (
	"fmt"
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

func (s *ModuleService) Delete(name string, courseName string) error {
	coursePath, err := pkg.GetPath(courseName, "./courses")
	if err != nil {
		return err
	}

	dirPath, err := pkg.GetPath(name, coursePath)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(dirPath); os.IsNotExist(err) {
		fmt.Println("Problem after RemoveAll")
		return err
	}

	return s.repo.Delete(name, courseName)
}
