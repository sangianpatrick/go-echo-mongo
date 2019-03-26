package model

import (
	"time"
)

// User is property info
type User struct {
	ID        string    `bson:"id"`
	FirstName string    `bson:"firstName"`
	LastName  string    `bson:"lastName"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

// Users is a slice of "User"
type Users []User
