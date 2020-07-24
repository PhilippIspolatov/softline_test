package main

import (
	"log"

	"github.com/PhilippIspolatov/softline_test/internal/db"
	userDelivery "github.com/PhilippIspolatov/softline_test/internal/user/delivery"
	userRepository "github.com/PhilippIspolatov/softline_test/internal/user/repository"
	userUseCase "github.com/PhilippIspolatov/softline_test/internal/user/usecase"

	sessionDelivery "github.com/PhilippIspolatov/softline_test/internal/session/delivery"
	sessionRepository "github.com/PhilippIspolatov/softline_test/internal/session/repository"
	sessionUseCase "github.com/PhilippIspolatov/softline_test/internal/session/usecase"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main()  {
	e := echo.New()

	db, err := db.NewDBConn("./db.json")

	if err != nil {
		return
	}

	defer func() {
		err := db.Conn.Close

		if err != nil {
			return
		}
	}()

	userRep := userRepository.NewUserRepository(db.Conn)
	sessionRep := sessionRepository.NewSessionRepository(db.Conn)

	userUCase := userUseCase.NewUserUseCase(userRep)
	sessionUCase := sessionUseCase.NewSessionUseCase(sessionRep, userRep)

	_ = userDelivery.NewUserDelivery(e, userUCase)
	_ = sessionDelivery.NewSessionDelivery(e, sessionUCase)

	log.Fatal(e.Start(":5000"))
}