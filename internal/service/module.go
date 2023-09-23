package service

import (
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
	path := pkg.GetPath(courseName, "./courses")

	dbFolderName, err := pkg.CreateUniqueFolder("module", path)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, courseName, dbFolderName)
}
