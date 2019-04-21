package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sangianpatrick/go-echo-mongo/helpers/wrapper"
)

// AppMiddleware is package that contains function for filtering request
type AppMiddleware struct {
	appName string
}

// CORS is a function that will filter the incoming request
func (am *AppMiddleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		contentType := c.Request().Header.Get("Content-Type")
		c.Response().Header().Set("Server", am.appName)
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods",
			"GET,PUT,POST,DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers",
			"Origin, X-Requested-With, Content-Type, Accept")

		c.Response().Header().Set("Accept", "application/json")
		if contentType != "application/json" {
			fmt.Println(contentType)
			return wrapper.Error(http.StatusNotAcceptable, "request is not acceptable due to policy", c)
		}
		if c.Request().Method == "OPTIONS" {
			return c.String(http.StatusOK, "")
		}
		return next(c)
	}
}

// InitAppMiddleware is a function that act as AppMiddleware constructor
func InitAppMiddleware(appName string) *AppMiddleware {
	return &AppMiddleware{
		appName: appName,
	}
}
