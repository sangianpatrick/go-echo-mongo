package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/sangianpatrick/go-echo-mongo/src/modules/user/handler"

	db "github.com/sangianpatrick/go-echo-mongo/helpers/database"
	"github.com/sangianpatrick/go-echo-mongo/src/modules/user/repository"
)

func init() {
	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	var mongoCredential = map[string]string{
		"host":     viper.GetString(`database.mongodb.host`),
		"user":     viper.GetString(`database.mongodb.user`),
		"password": viper.GetString(`database.mongodb.password`),
		"db":       viper.GetString(`database.mongodb.db`),
	}

	mongodb, err := db.GetMongoDB(mongoCredential)
	if err != nil {
		fmt.Println("MongoDB Error:", err)
		os.Exit(1)
	}

	defer mongodb.Logout()
	e := echo.New()

	urMongo := repository.NewUserRepositoryMongo(mongodb, "user")

	handler.NewUserHandler(e, urMongo)

	e.Start("localhost:8000")

}
