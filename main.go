package main

import (
	"fmt"
	"time"

	"github.com/sangianpatrick/go-mongo/config"
	"github.com/sangianpatrick/go-mongo/src/modules/user/model"
	"github.com/sangianpatrick/go-mongo/src/modules/user/repository"
)

func main() {
	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	defer db.Logout()

	urMongo := repository.NewUserRepositoryMongo(db, "user")

	saveUser(urMongo)

}

func saveUser(ur repository.UserRepository) {
	var u model.User

	u.ID = "User001"
	u.FirstName = "Patrick"
	u.LastName = "Maurits"
	u.Email = "patrickmaurits@gmail.com"
	u.Password = "14qwafzx"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	err := ur.Save(&u)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user has been succesfuly created")
	}
}
