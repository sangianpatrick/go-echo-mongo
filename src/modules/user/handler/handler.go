package handler

import (
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
		wrapper.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusNotFound, wrapper.ResponseError)
	}
	wrapper.ResponseSuccess["data"] = user
	wrapper.ResponseSuccess["message"] = "Detail of user."
	return c.JSON(http.StatusOK, wrapper.ResponseSuccess)
}

// GetAllUser is a function to return list of user
func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.uUcase.FindAll()
	if err != nil {
		wrapper.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusNotFound, wrapper.ResponseError)
	}
	if len(users) < 1 {
		wrapper.ResponseSuccess["message"] = "User list is empty"
		return c.JSON(http.StatusNoContent, wrapper.ResponseSuccess)
	}
	wrapper.ResponseSuccess["data"] = users
	wrapper.ResponseSuccess["message"] = "List of user"
	return c.JSON(http.StatusOK, wrapper.ResponseSuccess)
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
		wrapper.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, wrapper.ResponseError)
	}
	if ok, err := isRequestValid(&user); !ok {
		wrapper.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, wrapper.ResponseError)
	}

	err = h.uUcase.Save(&user)
	if err != nil {
		wrapper.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusInternalServerError, wrapper.ResponseError)
	}
	wrapper.ResponseSuccess["data"] = user
	wrapper.ResponseSuccess["message"] = "A user has successfuly created."
	return c.JSON(http.StatusCreated, wrapper.ResponseSuccess)
}
