package core

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model

	Name        string `gorm:"not null; unique"`
	Description string
	NameFolder  string

	Modules []Module `gorm:"many2many:courses_modules;"`
}

type Module struct {
	gorm.Model

	Name        string
	Description string
	NameFolder  string `gorm:"not null; unique"`

	Lessons []Lesson `gorm:"many2many:modules_lessons;"`
}

type Lesson struct {
	gorm.Model

	Name        string `gorm:"not null; unique"`
	Description string
	NameFolder  string
}
