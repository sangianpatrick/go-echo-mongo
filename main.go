package main

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/sangianpatrick/go-mongo/src/modules/user/handler"

	"github.com/sangianpatrick/go-mongo/config"
	"github.com/sangianpatrick/go-mongo/src/modules/user/repository"
)

func main() {
	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Logout()
	e := echo.New()

	urMongo := repository.NewUserRepositoryMongo(db, "user")

	handler.NewUserHandler(e, urMongo)

	e.Start("localhost:8000")

}
