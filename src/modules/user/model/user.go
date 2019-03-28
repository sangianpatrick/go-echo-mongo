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
	Password  string    `bson:"password" json:"password" validate:"required"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

//UserPublic is property that used for a response
type UserPublic struct {
	UserID    string    `bson:"userId" json:"userId"`
	FirstName string    `bson:"firstName" json:"firstName"`
	LastName  string    `bson:"lastName" json:"lastName"`
	Email     string    `bson:"email" json:"email" validate:"required,email"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
}

// Users is a slice of "User"
type Users []UserPublic
