package repository

import (
	"github.com/PhilippIspolatov/softline_test/internal/models"
	"github.com/PhilippIspolatov/softline_test/internal/tools"
	"github.com/PhilippIspolatov/softline_test/internal/user"
	"github.com/jackc/pgx"
)

type UserRepository struct {
	db *pgx.ConnPool
}

func NewUserRepository(db *pgx.ConnPool) user.Repository {
	return &UserRepository {
		db: db,
	}
}

func (ur *UserRepository) Insert(user *models.User) error {
	if err := ur.db.QueryRow("INSERT INTO users (nickname, email, password, phone) " +
		"VALUES ($1, $2, $3, $4) RETURNING nickname",
		user.Nickname,
		user.Email,
		user.Password,
		user.Phone).Scan(&user.Nickname); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Get(nickname string) (*models.User, error) {
	User := &models.User{}

	if err := ur.db.QueryRow("SELECT * FROM USERS WHERE nickname = $1", nickname).Scan(
		&User.Nickname, &User.Email, &User.Password, &User.Phone); err != nil {
		return nil, tools.UserDoesNotExist
	}

	return User, nil
}