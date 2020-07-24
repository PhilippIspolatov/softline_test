package usecase

import (
	"github.com/PhilippIspolatov/softline_test/internal/models"
	"github.com/PhilippIspolatov/softline_test/internal/tools"
	"github.com/PhilippIspolatov/softline_test/internal/user"
)

type UserUseCase struct {
	userRepository user.Repository
}

func NewUserUseCase(ur user.Repository) user.UseCase {
	return &UserUseCase{
		userRepository: ur,
	}
}

func (uUC *UserUseCase) CreateUser(user *models.User) error {
	_, err := uUC.GetUser(user.Nickname)

	if err == nil {
		return tools.AlreadyExist
	}

	user.PasswordHash()

	err = uUC.userRepository.Insert(user)

	if err != nil {
		return tools.ErrorCreatingUser
	}

	return nil
}

func (uUC *UserUseCase) GetUser(nickname string) (*models.User, error) {
	u, err := uUC.userRepository.Get(nickname)

	if err != nil {
		return nil, err
	}

	return u, nil
}