package user

import "github.com/PhilippIspolatov/softline_test/internal/models"

type Repository interface {
	Insert (user *models.User) error
	Get (nickname string) (*models.User, error)
}
