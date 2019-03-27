package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sangianpatrick/go-mongo/src/modules/user/model"
	"github.com/sangianpatrick/go-mongo/src/modules/user/repository"
	validator "gopkg.in/go-playground/validator.v9"
)

//ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

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

//GetUser function to get message
func (h *UserHandler) GetUser(c echo.Context) error {
	userID := c.Param("userID")
	user, err := h.uUcase.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

// GetAllUser is a function to return list of user
func (h *UserHandler) GetAllUser(c echo.Context) error {
	users, err := h.uUcase.FindAll()
	if err != nil {
		return c.JSON(http.StatusNotFound, ResponseError{Message: err.Error()})
	}
	if len(users) < 1 {
		return c.JSON(http.StatusNoContent, ResponseError{Message: "User list is empty"})
	}
	return c.JSON(http.StatusOK, users)
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
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	if ok, err := isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.uUcase.Save(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

// func (a *HttpArticleHandler) FetchArticle(c echo.Context) error {

// 	numS := c.QueryParam("num")
// 	num, _ := strconv.Atoi(numS)
// 	cursor := c.QueryParam("cursor")
// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}
// 	listAr, nextCursor, err := a.AUsecase.Fetch(ctx, cursor, int64(num))

// 	if err != nil {
// 		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
// 	}
// 	c.Response().Header().Set(`X-Cursor`, nextCursor)
// 	return c.JSON(http.StatusOK, listAr)
// }
