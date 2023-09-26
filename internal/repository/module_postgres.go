package repository

import (
	"errors"
	"fmt"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type ModulePostgres struct {
	db *gorm.DB
}

func NewModulePostgres(db *gorm.DB) *ModulePostgres {
	return &ModulePostgres{db: db}
}

func (r *ModulePostgres) Add(name string, description string, courseName string, folderName string) error {
	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	newModule := core.Module{
		Name:        name,
		Description: description,
		NameFolder:  folderName,
		CourseID:    course.ID,
	}

	if result := r.db.Create(&newModule); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ModulePostgres) Delete(id uint) (string, error) {
	var module core.Module
	if result := r.db.Where("id = ?", id).First(&module); result.Error != nil {
		return "", result.Error
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&module); result.Error != nil {
		return "", result.Error
	}

	return module.NameFolder, nil
}

func (r *ModulePostgres) Get(path string) (core.ModLes, error) {
	var modles core.ModLes

	var module core.Module

	if result := r.db.Where("name_folder = ?", path).First(&module); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return modles, fmt.Errorf("module not found for path: %s", path)
		}
		return modles, result.Error
	}

	var lessons []core.Lesson

	if result := r.db.Where("module_id = ?", module.ID).Find(&lessons); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return modles, fmt.Errorf("lessons not found for module ID: %d", module.ID)
		}
		return modles, result.Error
	}

	modles = core.ModLes{
		Module:  module,
		Lessons: lessons,
	}

	return modles, nil
}
