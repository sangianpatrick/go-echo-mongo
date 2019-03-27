package repository

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/sangianpatrick/go-echo-mongo/src/modules/user/model"
)

// userRepositoryMongo is an instance connection of MongoDB
type userRepositoryMongo struct {
	db         *mgo.Database
	collection string
}

// NewUserRepositoryMongo is function that will return userRepositoryMongo
func NewUserRepositoryMongo(db *mgo.Database, collection string) *userRepositoryMongo {
	return &userRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

// Save is a function to create new user
func (r *userRepositoryMongo) Save(user *model.User) error {
	err := r.db.C(r.collection).Insert(user)
	return err
}

// Update is a function to update existing user
func (r *userRepositoryMongo) Update(userID string, user *model.User) error {
	user.UpdatedAt = time.Now()
	err := r.db.C(r.collection).Update(bson.M{"userId": userID}, user)
	return err
}

// Delete is a function to remove a record from User list
func (r *userRepositoryMongo) Delete(userID string) error {
	err := r.db.C(r.collection).Remove(bson.M{"userId": userID})
	return err
}

// FindByID is a function to get one user by ID
func (r *userRepositoryMongo) FindByID(userID string) (*model.User, error) {
	var user model.User
	err := r.db.C(r.collection).Find(bson.M{"userId": userID}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll is a function to get User list
func (r *userRepositoryMongo) FindAll() (model.Users, error) {
	var users model.Users
	err := r.db.C(r.collection).Find(bson.M{}).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
