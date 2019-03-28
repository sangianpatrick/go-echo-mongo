package main

import (
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	mdl "github.com/sangianpatrick/go-echo-mongo/middleware"
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
	appName := viper.GetString(`app.name`)
	appHost := viper.GetString(`app.domain`)
	appPort := viper.GetString(`app.port`)

	mongodb, err := db.GetMongoDB(mongoCredential)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer mongodb.Logout()

	e := echo.New()
	appMiddleware := mdl.InitAppMiddleware(appName)
	e.Use(appMiddleware.CORS)

	urMongo := repository.NewUserRepositoryMongo(mongodb, "user")

	handler.NewUserHandler(e, urMongo)

	e.Start(fmt.Sprintf(`%s:%s`, appHost, appPort))

}
