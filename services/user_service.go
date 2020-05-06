package services

import (
	"fmt"

	"github.com/sampado/bookstore_users-api/domain/users"
)

// Gets a user from the BBDD.
func GetUser() {}

func CreateUser(user *users.User) (*users.User, error) {
	fmt.Println("User created!")
	return user, nil
}

func FindUser() {}
