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

func (r *ModulePostgres) Add(name string, description string, orderID uint, courseName string) error {
	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	newModule := core.Module{
		Name:        name,
		Description: description,
		CourseID:    course.ID,
		OrderID:     orderID - 1,
	}

	if err := r.db.Where("name = ? AND course_id = ?", newModule.Name, newModule.CourseID).First(&core.Module{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if result := r.db.Create(&newModule); result.Error != nil {
				return result.Error
			}

			return nil
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("this module already exists")
	}
}

func (r *ModulePostgres) Delete(id uint) ([]uint, error) {
	var module core.Module
	var lessons []core.Lesson
	var lessonsID []uint

	if result := r.db.Where("id = ?", id).First(&module); result.Error != nil {
		return nil, result.Error
	}

	if result := r.db.Where("module_id = ?", module.ID).Find(&lessons); result.Error != nil {
		return nil, result.Error
	}

	for _, lesson := range lessons {
		lessonsID = append(lessonsID, lesson.ID)
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&module); result.Error != nil {
		return nil, result.Error
	}

	return lessonsID, nil
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
