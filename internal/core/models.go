package core

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model

	Name        string `gorm:"not null; unique"`
	Description string
	NameFolder  string `gorm:"not null"`
}

type Module struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string
	NameFolder  string `gorm:"not null"`
	CourseID    uint

	Course Course `gorm:"foreignKey:CourseID"`
}

type Lesson struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string
	NameFile    string `gorm:"not null"`
	ModuleID    uint

	Module Module `gorm:"foreignKey:ModuleID"`
}