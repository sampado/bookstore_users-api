package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sampado/bookstore_users-api/domain/users"
	"github.com/sampado/bookstore_users-api/services"
	"github.com/sampado/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(userErr.Error())
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)
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
		fmt.Println(saveErr)
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(userErr.Error())
		c.JSON(err.Status, err)
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		error := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(http.StatusBadRequest, error)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
