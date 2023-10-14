package core

import (
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type Course struct {
	GormModel

	Name        string `gorm:"not null; unique"`
	Description string `gorm:"not null"`
}

type Module struct {
	GormModel

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	OrderID     uint   `gorm:"not null"`
	CourseID    uint

	Course Course `gorm:"foreignKey:CourseID" json:"-"`
}

type Lesson struct {
	GormModel

	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	OrderID     uint   `gorm:"not null"`
	ModuleID    uint

	Module Module `gorm:"foreignKey:ModuleID" json:"-"`
}

type Sessions struct {
	GormModel
	RefreshToken   string
	ExpirationTime time.Time
}

type Email struct {
	GormModel

	Address string `gorm:"not null; unique"`
}
