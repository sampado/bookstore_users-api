package services

import (
	"github.com/sampado/bookstore_users-api/domain/users"
	dateutils "github.com/sampado/bookstore_users-api/utils/date_utils"
	"github.com/sampado/bookstore_users-api/utils/errors"
)

// GetUser is a service to get a user from the BBDD.
func GetUser(userID int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userID}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser is a service to create a user and persist it into the BBDD.
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = dateutils.GetNowDBFormat()
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser used to update an user in the BBDD
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		// full update
		if err := user.Validate(); err != nil {
			return nil, err
		}
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

// DeleteUser is used to delete a given user
func DeleteUser(userID int64) *errors.RestError {
	user := &users.User{Id: userID}
	return user.Delete()
}

// FindUserByStatus is used to find users
func FindUserByStatus(status string) ([]users.User, *errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
