package services

import (
	"github.com/sampado/bookstore_users-api/domain/users"
	cryptoutils "github.com/sampado/bookstore_users-api/utils/crypto_utils"
	dateutils "github.com/sampado/bookstore_users-api/utils/date_utils"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {
}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *rest_errors.RestError)
	CreateUser(users.User) (*users.User, *rest_errors.RestError)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestError)
	DeleteUser(int64) *rest_errors.RestError
	FindUserByStatus(string) (users.Users, *rest_errors.RestError)
	LoginUser(users.LoginRequest) (*users.User, *rest_errors.RestError)
}

// GetUser is a service to get a user from the BBDD.
func (s *usersService) GetUser(userID int64) (*users.User, *rest_errors.RestError) {
	user := &users.User{Id: userID}
	err := user.Get()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser is a service to create a user and persist it into the BBDD.
func (s *usersService) CreateUser(user users.User) (*users.User, *rest_errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = dateutils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = cryptoutils.GetMd5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser used to update an user in the BBDD
func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestError) {
	current, err := s.GetUser(user.Id)
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
func (s *usersService) DeleteUser(userID int64) *rest_errors.RestError {
	user := &users.User{Id: userID}
	return user.Delete()
}

// FindUserByStatus is used to find users
func (s *usersService) FindUserByStatus(status string) (users.Users, *rest_errors.RestError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *rest_errors.RestError) {
	dao := &users.User{
		Email:    request.Email,
		Password: cryptoutils.GetMd5(request.Password),
	}

	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
