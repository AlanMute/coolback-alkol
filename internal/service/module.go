package service

import (
	"fmt"
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

func (s *ModuleService) Add(name string, description string, orderID uint, courseID uint) error {
	if name == "" {
		name = "New Module" + fmt.Sprint(orderID)
	}

	if err := s.repo.Add(name, description, orderID, courseID); err != nil {
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
		fileName := strconv.FormatUint(uint64(lessonToDelete), 10) + ext

		pathToLesson := filepath.Join("./lessons", fileName)

		if err := os.Remove(pathToLesson); err != nil {
			return err
		}
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
