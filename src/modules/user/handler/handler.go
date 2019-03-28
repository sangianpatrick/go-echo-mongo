package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	wp "github.com/sangianpatrick/go-echo-mongo/helpers/wrapper"
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
		wp.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusNotFound, wp.ResponseError)
	}
	wp.ResponseSuccess["data"] = user
	wp.ResponseSuccess["message"] = "Detail of user."
	return c.JSON(http.StatusOK, wp.ResponseSuccess)
}

// GetAllUser is a function to return list of user
func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.uUcase.FindAll()
	if err != nil {
		wp.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusNotFound, wp.ResponseError)
	}
	if len(users) < 1 {
		wp.ResponseSuccess["message"] = "User list is empty"
		return c.JSON(http.StatusNoContent, wp.ResponseSuccess)
	}
	wp.ResponseSuccess["data"] = users
	wp.ResponseSuccess["message"] = "List of user"
	return c.JSON(http.StatusOK, wp.ResponseSuccess)
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
		wp.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusUnprocessableEntity, wp.ResponseError)
	}
	if ok, err := isRequestValid(&user); !ok {
		wp.ResponseError["message"] = err.Error()
		return c.JSON(http.StatusBadRequest, wp.ResponseError)
	}

	err = h.uUcase.Save(&user)
	if err != nil {
		wp.ResponseError["message"] = "User is already created."
		return c.JSON(http.StatusConflict, wp.ResponseError)
	}
	wp.ResponseSuccess["message"] = "A user has successfuly created."
	return c.JSON(http.StatusCreated, wp.ResponseSuccess)
}
