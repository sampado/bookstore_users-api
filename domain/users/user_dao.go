package users

import (
	"fmt"

	"github.com/sampado/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.DataCreated = result.DataCreated
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName

	return nil // no errors
}

func (user *User) Save() *errors.RestError {
	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exits", user.Id))
	}

	usersDB[user.Id] = user
	return nil // no error
}
