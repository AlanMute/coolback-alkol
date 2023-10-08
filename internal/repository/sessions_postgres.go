package repository

import (
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/jinzhu/gorm"
)

type SessionPostgres struct {
	db                   *gorm.DB
	deleteRoutineStarted bool
}

func NewSessionPostgres(db *gorm.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) Add(session core.Sessions) error {
	if result := r.db.Create(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SessionPostgres) CheckRefresh(token string) error {
	var session core.Sessions

	if !r.deleteRoutineStarted {
		go r.deleteExpiredTokens()
		r.deleteRoutineStarted = true
	}

	if result := r.db.Where("refresh_token = ? AND expiration_time > ?", token, time.Now()).First(&session); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SessionPostgres) deleteExpiredTokens() {
	for {
		r.db.Where("expiration_time < ?", time.Now()).Delete(&core.Sessions{})

		time.Sleep(24 * time.Hour)
	}
}
