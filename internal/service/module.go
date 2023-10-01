package service

import (
	"os"

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

func (s *ModuleService) Get(id int) (core.ModLes, error) {
	modles, err := s.repo.Get(id)
	if err != nil {
		return modles, err
	}

	return modles, nil
}

func (s *ModuleService) Put(id int, name string, desc string, orderid uint) error {
	err := s.repo.Put(id, name, desc, orderid)

	if err != nil {
		return err
	}

	return nil
}
