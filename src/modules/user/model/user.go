package model

import (
	"time"
)

// User is property info
type User struct {
	UserID    string    `bson:"userId" json:"userId" validate:"required"`
	FirstName string    `bson:"firstName" json:"firstName" validate:"required"`
	LastName  string    `bson:"lastName" json:"lastName" validate:"required"`
	Email     string    `bson:"email" json:"email" validate:"required,email"`
	Password  string    `bson:"password" json:"-" validate:"required"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// Users is a slice of "User"
type Users []User
