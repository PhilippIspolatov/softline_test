package repository

import (
	"github.com/PhilippIspolatov/softline_test/internal/models"
	"github.com/jackc/pgx"
)

type SessionRepository struct {
	db *pgx.ConnPool
}

func NewSessionRepository(db *pgx.ConnPool) *SessionRepository {
	return &SessionRepository{
		db:db,
		}
}

func (sr *SessionRepository) Create(session *models.Session) error {
	if err := sr.db.QueryRow("INSERT INTO sessions VALUES ($1, $2, $3) RETURNING nickname",
		session.Nickname, session.Cookie, session.Expiration).Scan(&session.Nickname); err != nil {
		return err
	}

	return nil
}

func (sr *SessionRepository) GetByNickname(nickname string) (*models.Session, error) {
	s := &models.Session{}
	if err := sr.db.QueryRow("SELECT * FROM sessions WHERE nickname = $1", nickname).Scan(
		&s.Nickname,
		&s.Cookie,
		&s.Expiration); err != nil {
		return nil, err
	}

	return s, nil
}

func (sr *SessionRepository) GetByCookie(cookie string) (*models.Session, error) {
	s := &models.Session{}

	if err := sr.db.QueryRow("SELECT * FROM sessions WHERE cookie = $1", cookie).Scan(
		&s.Nickname,
		&s.Cookie,
		&s.Expiration); err != nil {
		return nil, err
	}

	return s, nil
}