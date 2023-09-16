package service

import "github.com/KrizzMU/coolback-alkol/internal/repository"

type ModuleService struct {
	repo repository.Module
}

func NewModuleService(repo repository.Module) *ModuleService {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) Add(name string, description string, id int, folderName string) error {

	return s.repo.Add(name, description, id, folderName)
}
