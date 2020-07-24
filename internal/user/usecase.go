package user

import "github.com/PhilippIspolatov/softline_test/internal/models"

type UseCase interface {
	CreateUser(user *models.User) error
	GetUser (nickname string) (*models.User, error)
}