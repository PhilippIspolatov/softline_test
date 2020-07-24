package main

import (
	"log"

	"github.com/PhilippIspolatov/softline_test/internal/db"
	userDelivery "github.com/PhilippIspolatov/softline_test/internal/user/delivery"
	userRepository "github.com/PhilippIspolatov/softline_test/internal/user/repository"
	userUseCase "github.com/PhilippIspolatov/softline_test/internal/user/usecase"
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
	userUCase := userUseCase.NewUserUseCase(userRep)
	_ = userDelivery.NewUserDelivery(e, userUCase)

	log.Fatal(e.Start(":5000"))
}