package users

import (
	"strings"

	"github.com/sampado/bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:id`
	FirstName   string `json:firstName`
	LastName    string `json:lastName`
	Email       string `json:email`
	DataCreated string `json:dataCreated`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
