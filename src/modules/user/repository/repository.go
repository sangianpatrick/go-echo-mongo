package repository

import (
	"github.com/sangianpatrick/go-mongo/src/modules/user/model"
)

// UserRepository interface is a list of method of user's model
type UserRepository interface {
	Save(*model.User) error
	Update(string, *model.User) error
	Delete(string) error
	FindByID(string) (*model.User, error)
	FindAll() (model.Users, error)
}
