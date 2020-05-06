package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sampado/bookstore_users-api/domain/users"
	"github.com/sampado/bookstore_users-api/services"
	"github.com/sampado/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User

	// long way
	// ---------
	// body, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// TODO: handle the error
	// 	return
	// }
	// if err := json.Unmarshal(body, &user); err != nil {
	// 	// TODO: handle the error
	// 	fmt.Println(err.Error)
	// 	return
	// }

	// gin framework way
	// -------
	if err := c.ShouldBindJSON(&user); err != nil {
		error := errors.NewBadRequestError(err.Error())
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, error)
		return
	}

	fmt.Print("User")
	fmt.Println(user)

	result, saveErr := services.CreateUser(&user)
	if saveErr != nil {
		error := errors.RestError{
			Message: "Error: Unable create an User",
			Status:  http.StatusInternalServerError,
			Error:   http.StatusText(http.StatusInternalServerError),
		}
		fmt.Println(error)
		c.JSON(http.StatusBadRequest, error)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
