package wrapper

import (
	"github.com/labstack/echo"
)

// Props contains properties to be responded
type Props struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

// Data returns wrapped success data
func Data(code int, data interface{}, message string, c echo.Context) error {
	props := &Props{
		Code:    code,
		Data:    data,
		Message: message,
		Success: true,
	}
	return c.JSON(code, props)
}

func Error(code int, message string, c echo.Context) error {
	props := &Props{
		Code:    code,
		Data:    nil,
		Message: message,
		Success: false,
	}
	return c.JSON(code, props)
}
