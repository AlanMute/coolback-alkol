package repository

import (
	"errors"
	"fmt"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type CoursePostgres struct {
	db *gorm.DB
}

func NewCoursePostgres(db *gorm.DB) *CoursePostgres {
	return &CoursePostgres{db: db}
}

func (r *CoursePostgres) Add(name string, description string) error {
	newCourse := core.Course{
		Name:        name,
		Description: description,
	}

	if result := r.db.Create(&newCourse); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CoursePostgres) Delete(id uint) ([]uint, error) {
	var course core.Course
	var modules []core.Module
	var lessonsID []uint

	if result := r.db.Where("id = ?", id).First(&course); result.Error != nil {
		return nil, result.Error
	}

	if result := r.db.Where("course_id = ?", id).Find(&modules); result.Error != nil {
		return nil, result.Error
	}

	for _, module := range modules {
		var lessons []core.Lesson

		if result := r.db.Where("module_id = ?", module.ID).Find(&lessons); result.Error != nil {
			return nil, result.Error
		}

		for _, lesson := range lessons {
			lessonsID = append(lessonsID, lesson.ID)
		}
	}

	if result := r.db.Where("id = ?", id).Unscoped().Delete(&course); result.Error != nil {
		return nil, result.Error
	}

	fmt.Println(lessonsID)

	return lessonsID, nil
}

func (r *CoursePostgres) GetByName(name string) ([]core.Course, error) {

	var courses []core.Course
	r.db.Where("name ILIKE ?", "%"+name+"%").Find(&courses)

	return courses, nil
}

func (r *CoursePostgres) GetAll() ([]core.Course, error) {

	var courses []core.Course
	r.db.Find(&courses)
	return courses, nil
}

func (r *CoursePostgres) Get(id int) (core.CourseContent, error) {

	var content core.CourseContent

	var course core.Course

	if result := r.db.Where("id = ?", id).Find(&course); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return content, fmt.Errorf("Course not find for id: %d", id)
		}
		return content, result.Error
	}

	var modles []core.ModLes
	var modules []core.Module

	if result := r.db.Where("course_id = ?", id).Order("order_id ASC").Find(&modules); result.Error != nil {
		return content, result.Error
	}

	for _, m := range modules {

		var lessons []core.Lesson

		if result := r.db.Where("module_id = ?", m.ID).Order("order_id ASC").Find(&lessons); result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return content, fmt.Errorf("lessons not found for module ID: %d", m.ID)
			}
			return content, result.Error
		}

		modles = append(modles, core.ModLes{Module: m, Lessons: lessons})
	}

	content = core.CourseContent{
		Course:  course,
		Modules: modles,
	}

	return content, nil
}

func (r *CoursePostgres) Put(id int, name string, desc string) error {
	var course core.Course

	if result := r.db.Where("id = ?", id).Find(&course); result.Error != nil {
		return result.Error
	}

	if name != "" {
		course.Name = name
	}
	if desc != "" {
		course.Description = desc
	}

	if result := r.db.Save(&course); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *CoursePostgres) CheckID(id uint) error {
	var course core.Course

	err := r.db.Where("id = ?", id).First(&course).Error

	if gorm.IsRecordNotFoundError(err) {
		return fmt.Errorf("no record with such ID")
	} else if err != nil {
		return err
	}

	return nil
}
