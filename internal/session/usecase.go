package session

import "github.com/PhilippIspolatov/softline_test/internal/models"


type UseCase interface {
	LogIn(nickname string, password string) (*models.Session, error)
	LogInByCookie(cookie string) (*models.Session, error)
}