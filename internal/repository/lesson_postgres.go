package repository

import (
	"errors"
	"fmt"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type LessonPostgres struct {
	db *gorm.DB
}

func NewLessonPostgres(db *gorm.DB) *LessonPostgres {
	return &LessonPostgres{db: db}
}

func (r *LessonPostgres) Add(name string, description string, fileName string, courseName string, moduleName string) error {
	var course core.Course
	if result := r.db.Where("name = ?", courseName).First(&course); result.Error != nil {
		return result.Error
	}

	var module core.Module
	if result := r.db.Where("name = ? AND course_id = ?", moduleName, course.ID).First(&module); result.Error != nil {
		return result.Error
	}

	newLesson := core.Lesson{
		Name:        name,
		Description: description,
		NameFile:    fileName,
		ModuleID:    module.ID,
	}

	if result := r.db.Create(&newLesson); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *LessonPostgres) Delete(id uint) (string, error) {
	var lesson core.Lesson
	if result := r.db.Where("id = ?", id).First(&lesson).Unscoped().Delete(&lesson); result.Error != nil {
		return "", result.Error
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&lesson); result.Error != nil {
		return "", result.Error
	}

	return lesson.NameFile, nil
}

func (r *LessonPostgres) Get(moduleid int, orderid int) (core.Lesson, error) {

	var lesson core.Lesson

	if result := r.db.Where("module_id = ? AND order_id = ?", moduleid, orderid).Find(&lesson); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return lesson, fmt.Errorf("Lesson not found for module_id = %d AND order_id = %d", moduleid, orderid)
		}
		return lesson, result.Error
	}

	return lesson, nil
}

func (r *LessonPostgres) Put(id int, name string, desc string, orderID uint) error {

	var lesson core.Lesson

	if err := r.db.Where("id = ?", id).Find(&lesson).Error; err != nil {
		return err
	}

	if name != "" {
		lesson.Name = name
	}
	if desc != "" {
		lesson.Description = desc
	}
	if orderID != 0 {
		lesson.OrderID = orderID
	}

	if err := r.db.Save(&lesson).Error; err != nil {
		return err
	}

	return nil
}
