package session

import "github.com/PhilippIspolatov/softline_test/internal/models"

type Repository interface {
	Create(session *models.Session) error
	GetByNickname(nickname string) (*models.Session, error)
	GetByCookie (cookie string) (*models.Session, error)
}
