package users

import (
	"fmt"
	"log"

	"github.com/sampado/bookstore_users-api/datasources/mysql/users_db"
	dateutils "github.com/sampado/bookstore_users-api/utils/date_utils"
	"github.com/sampado/bookstore_users-api/utils/errors"
	mysqlutils "github.com/sampado/bookstore_users-api/utils/mysql"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser = "DELETE FROM users WHERE id=?"
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		log.Println("error preparing the statement")
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // this will be excecuted right before the function ends/ on any of the return statements

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DataCreated); err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil // no errors
}

func (user *User) Save() *errors.RestError {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		log.Println("error preparing the statement")
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // this will be excecuted right before the function ends/ on any of the return statements

	user.DataCreated = dateutils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DataCreated)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get the last user ID: %s", err.Error()))
	}

	user.Id = userId

	return nil // no error
}

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysqlutils.ParseError(err)
	}

	return nil
}
