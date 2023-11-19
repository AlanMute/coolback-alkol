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

func (r *ModulePostgres) Add(name string, description string, orderID uint, courseID uint) error {
	newModule := core.Module{
		Name:        name,
		Description: description,
		CourseID:    courseID,
		OrderID:     orderID,
	}

	if err := r.db.Where("name = ? AND course_id = ? OR order_id = ? AND course_id = ?",
		newModule.Name, newModule.CourseID, newModule.OrderID, newModule.CourseID).First(&core.Module{}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			if result := r.db.Create(&newModule); result.Error != nil {
				return result.Error
			}

			return nil
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("this module already exists or order_id is taken")
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

func (r *ModulePostgres) Get(id int) (core.ModLes, error) {
	var modles core.ModLes

	var module core.Module

	if result := r.db.Where("id = ?", id).First(&module); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return modles, fmt.Errorf("module not found for id: %d", id)
		}
		return modles, result.Error
	}

	var lessons []core.Lesson

	if result := r.db.Where("module_id = ?", module.ID).Order("order_id ASC").Find(&lessons); result.Error != nil {
		return modles, result.Error
	}

	modles = core.ModLes{
		Module:  module,
		Lessons: lessons,
	}

	return modles, nil
}

func (r *ModulePostgres) Put(id int, name string, desc string, orderid uint) error {
	var module core.Module

	if result := r.db.Where("id = ?", id).Find(&module); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("module not found for id: %d", id)
		}
		return result.Error
	}

	if name != "" {
		module.Name = name
	}
	if desc != "" {
		module.Description = desc
	}
	if orderid != 0 {
		module.OrderID = orderid
	}

	if result := r.db.Save(&module); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ModulePostgres) CheckID(id uint) error {
	var module core.Module

	err := r.db.Where("id = ?", id).First(&module).Error

	if gorm.IsRecordNotFoundError(err) {
		return fmt.Errorf("no record with such ID")
	} else if err != nil {
		return err
	}

	return nil
}
