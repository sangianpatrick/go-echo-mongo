package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sangianpatrick/go-echo-mongo/helpers/wrapper"
	"github.com/sangianpatrick/go-echo-mongo/src/modules/user/model"
	"github.com/sangianpatrick/go-echo-mongo/src/modules/user/repository"
	validator "gopkg.in/go-playground/validator.v9"
)

// UserHandler represent the httphandler for article
type UserHandler struct {
	uUcase repository.UserRepository
}

// NewUserHandler is constructor
func NewUserHandler(e *echo.Echo, ur repository.UserRepository) {
	uh := &UserHandler{
		uUcase: ur,
	}
	e.POST("/users", uh.CreateUser)
	e.GET("/users/:userID", uh.GetUser)
	e.GET("/users", uh.GetAllUser)

}

// GetUser function to get message
func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("userID")
	user, err := h.uUcase.FindByID(userID)
	if err != nil {
		fmt.Println(err)
		return wrapper.Error(http.StatusNotFound, err.Error(), c)
	}
	return wrapper.Data(http.StatusOK, user, "user detail", c)
}

// GetAllUser is a function to return list of user
func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.uUcase.FindAll()
	if err != nil {
		return wrapper.Error(http.StatusNotFound, err.Error(), c)
	}
	if len(users) < 1 {
		return wrapper.Data(http.StatusNoContent, nil, "no content", c)
	}
	return wrapper.Data(http.StatusOK, users, "list of user", c)
}

// isRequestValid is function that act as request body validator
func isRequestValid(m *model.User) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateUser is a function to store new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var user model.User
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err := c.Bind(&user)
	if err != nil {
		return wrapper.Error(http.StatusUnprocessableEntity, err.Error(), c)
	}
	if ok, err := isRequestValid(&user); !ok {
		return wrapper.Error(http.StatusBadRequest, err.Error(), c)
	}

	err = h.uUcase.Save(&user)
	if err != nil {
		fmt.Println(err.Error())
		return wrapper.Error(http.StatusConflict, "user is already created", c)
	}
	return wrapper.Data(http.StatusCreated, nil, "user is created", c)
}
