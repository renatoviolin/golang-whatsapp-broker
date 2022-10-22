package services

import (
	"errors"
	"os"
)

type TokenService struct {
	SavedToken string
}

func NewTokenService() TokenService {
	return TokenService{SavedToken: os.Getenv("TOKEN")}
}

func (h *TokenService) VerifyToken(token string) error {
	if token != h.SavedToken {
		return errors.New("invalid token")
	}

	return nil
}
