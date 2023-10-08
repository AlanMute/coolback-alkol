package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type SessionPostgres struct {
	db *gorm.DB
}

func NewSessionPostgres(db *gorm.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) Add(refreshToken string, t time.Time) error {
	session := core.Sessions{
		RefreshToken:   refreshToken,
		ExpirationTime: t,
	}

	if result := r.db.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}
