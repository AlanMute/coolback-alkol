package service

import (
	"errors"
	"os"
	"time"

	"github.com/KrizzMU/coolback-alkol/internal/core"
	"github.com/KrizzMU/coolback-alkol/pkg/auth"
)

type AdminService struct {
	tockenManager auth.TokenManager
}

func NewAdminService(t auth.TokenManager) *AdminService {
	return &AdminService{tockenManager: t}
}

func (s *AdminService) SignIn(login string, pass string) (core.Tokens, error) {
	if login != os.Getenv("ADM_LOGIN") || pass != os.Getenv("ADM_PASSWORD") {
		return core.Tokens{}, errors.New("Uncorrect login or password!")
	}

	accessToken, err := s.tockenManager.NewAccessToken("admin", time.Hour)

	if err != nil {
		return core.Tokens{}, err
	}
	refreshToken, err := s.tockenManager.NewRefreshToken()

	if err != nil {
		return core.Tokens{}, err
	}

	token := core.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return token, nil
}
