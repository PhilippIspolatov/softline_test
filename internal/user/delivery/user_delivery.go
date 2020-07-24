package delivery

import (
	"net/http"

	"github.com/PhilippIspolatov/softline_test/internal/models"
	"github.com/PhilippIspolatov/softline_test/internal/tools"
	"github.com/PhilippIspolatov/softline_test/internal/user"
	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUseCase user.UseCase
}

func NewUserDelivery(router *echo.Echo, uUC user.UseCase) *UserDelivery {
	ud := &UserDelivery{
		userUseCase: uUC,
	}

	router.POST("/api/v1/user/create", ud.CreateUser())

	return ud
}

func (ud *UserDelivery) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &models.User{}
		err := c.Bind(req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, tools.Error{
				Error: tools.HttpBadRequest.Error(),
			})
		}

		u, err := models.NewUser(req.Nickname, req.Email, req.Password, req.Phone)

		if err != nil {
			return c.JSON(http.StatusBadRequest, tools.Error{
				Error:tools.HttpBadRequest.Error(),
			})
		}

		err = ud.userUseCase.CreateUser(u)

		if err == tools.AlreadyExist {
			return c.JSON(http.StatusConflict, tools.Error{
				Error: tools.HttpConflict.Error(),
			})
		}
		if err != nil {
				return c.JSON(http.StatusBadRequest, tools.Error{
					Error:tools.HttpBadRequest.Error(),
				})
			}
		return c.JSON(http.StatusOK, tools.Message{
			Message: "success",
		})
	}
}