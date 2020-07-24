package delivery

import (
	"net/http"

	"github.com/PhilippIspolatov/softline_test/internal/session"
	"github.com/PhilippIspolatov/softline_test/internal/tools"
	"github.com/labstack/echo/v4"
)

type SessionDelivery struct {
	sessionUseCase session.UseCase
}

func NewSessionDelivery(router *echo.Echo, sUC session.UseCase) *SessionDelivery {
	sd := &SessionDelivery{
		sessionUseCase: sUC,
	}

	router.POST("/api/v1/user/login", sd.LogIn())

	return sd
}

func (sd *SessionDelivery) LogIn() echo.HandlerFunc {

	type req struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}

	return func(c echo.Context) error {
		req := &req{}

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, tools.Error{
				Error: tools.HttpBadRequest.Error(),
			})
		}

		oldCookie, err := c.Cookie("softline_test")
		if err == nil {
			_, err := sd.sessionUseCase.LogInByCookie(oldCookie.Value)
			if err == nil {
				return c.JSON(http.StatusConflict, tools.Error{
					Error: tools.HttpAlreadyAuthenticate.Error(),
				})
			}
		}

		s, err := sd.sessionUseCase.LogIn(req.Nickname, req.Password)

		if err != nil {
			return c.JSON(http.StatusBadRequest, tools.Error{
				Error: err.Error(),
			})
		}

		cookie := &http.Cookie{
			Name:       "softline_test",
			Value:      s.Cookie,
			MaxAge:     60 * 60 * 24,
			HttpOnly:   true,
			SameSite:   http.SameSiteNoneMode,

		}

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, tools.Message{Message: "success"})
	}
}