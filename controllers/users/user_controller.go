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

func Get(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
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

	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		fmt.Println(saveErr)
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		error := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(http.StatusBadRequest, error)
		return
	}

	user.Id = userID

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userID, userErr := getUserID(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if err := services.UsersService.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status, foundParam := c.GetQuery("status")
	if !foundParam {
		// define default value
		status = "active"
	}

	users, err := services.UsersService.FindUserByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func getUserID(userIdParam string) (int64, *errors.RestError) {
	userID, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError(userErr.Error())
		return -1, err
	}

	return userID, nil
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
