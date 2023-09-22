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

func (s *ModuleService) Add(name string, description string, id int) error {
	var path string // ex. "/course_1"

	folderName, err := pkg.UniqueFolder("module", path)
	if err != nil {
		return err
	}

	return s.repo.Add(name, description, id, folderName)
}
