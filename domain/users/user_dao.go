package users

import (
	"fmt"
	"strings"

	"github.com/sampado/bookstore_users-api/datasources/mysql/users_db"
	mysqlutils "github.com/sampado/bookstore_users-api/utils/mysql"
	"github.com/sampado/bookstore_utils-go/logger"
	"github.com/sampado/bookstore_utils-go/rest_errors"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?)"
	queryGetUser                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id = ?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?"
	queryFindUserByEmailAndPass = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email = ? and password = ?"
)

func (user *User) Get() *rest_errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error preparing get user statement", err)
		return rest_errors.NewInternalServerError("error preparing the statement", err)
	}
	defer stmt.Close() // this will be excecuted right before the function ends/ on any of the return statements

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil // no errors
}

func (user *User) Save() *rest_errors.RestError {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error preparing save user statement", err)
		return rest_errors.NewInternalServerError("error preparing save user statement", err)
	}
	defer stmt.Close() // this will be excecuted right before the function ends/ on any of the return statements

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get the last user ID: %s", err)
		return rest_errors.NewInternalServerError("error when trying to get the last user", err)
	}

	user.Id = userId

	return nil // no error
}

func (user *User) Update() *rest_errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *rest_errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return rest_errors.NewInternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *rest_errors.RestError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	defer rows.Close()
	users := make([]User, 0)

	for rows.Next() {
		var user User = User{}
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseError(err)
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return users, nil
}

func (user *User) FindByEmailAndPassword() *rest_errors.RestError {
	stmt, err := users_db.Client.Prepare(queryFindUserByEmailAndPass)
	if err != nil {
		logger.Error("error preparing find user by email and password statement", err)
		return rest_errors.NewInternalServerError("error preparing the statement", err)
	}
	defer stmt.Close() // this will be excecuted right before the function ends/ on any of the return statements

	result := stmt.QueryRow(user.Email, user.Password)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		if strings.Contains(err.Error(), mysqlutils.ErrorNoRows) {
			return rest_errors.NewNotFoundError("invalid user credentials")
		}
		logger.Error("error when trying to find user by email and password statement", err)
		return mysqlutils.ParseError(err)
	}

	return nil // no errors
}
