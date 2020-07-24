package usecase

import (
	"time"

	"github.com/PhilippIspolatov/softline_test/internal/models"
	"github.com/PhilippIspolatov/softline_test/internal/session"
	"github.com/PhilippIspolatov/softline_test/internal/tools"
	"github.com/PhilippIspolatov/softline_test/internal/user"
	"github.com/google/uuid"
)

type SessionUseCase struct {
	sessionRepository session.Repository
	userRepository user.Repository
}

func NewSessionUseCase(sr session.Repository, ur user.Repository) *SessionUseCase {
	return &SessionUseCase{
		sessionRepository: sr,
		userRepository: ur,
	}
}

func (sUC *SessionUseCase) LogIn(nickname string, password string) (*models.Session, error) {
	u, err := sUC.userRepository.Get(nickname)
	if err != nil {
		return nil, tools.UserDoesNotExist
	}

	if !u.CheckPassword(password) {
		return nil, tools.WrongPassword
	}

	s := &models.Session{
		Nickname:   nickname,
		Cookie:     uuid.New().String(),
		Expiration: time.Now().AddDate(0,0,1),
	}

	err = sUC.sessionRepository.Create(s)

	if err != nil {
		return nil, err
	}

	return s, err
}

func (sUC *SessionUseCase) LogInByCookie(cookie string) (*models.Session, error) {
	s, err := sUC.sessionRepository.GetByCookie(cookie)

	if err != nil {
		return nil, tools.UserDoesNotExist
	}

	return s, nil
}