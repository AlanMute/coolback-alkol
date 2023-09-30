package core

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model

	Name        string `gorm:"not null; unique"`
	Description string `gorm:"not null"`
	NameFolder  string `gorm:"not null"` // DELETE
}

type Module struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	NameFolder  string `gorm:"not null"` // DELETE
	OrderID     uint   `gorm:"not null; unique"`
	CourseID    uint

	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

type Lesson struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	NameFile    string `gorm:"not null"` // DELETE
	OrderID     uint   `gorm:"not null; unique"`
	ModuleID    uint

	Module Module `gorm:"foreignKey:ModuleID" json:"-"`
}
