package users

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"	
	"strconv"
	// "io/ioutil"
	// "encoding/json"
	"github.com/selvamshan/bookstore_user-api/domain/users"
	"github.com/selvamshan/bookstore_user-api/services"
	"github.com/selvamshan/bookstore_user-api/utils/errors"
)

func getUserId(userIdParam string)(int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be number")		
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User	
	
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {		
		c.JSON(saveErr.Status, saveErr)
		return
	}
	
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public")=="true"))
}


func GetUser(c *gin.Context) {	
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {		
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {		
		c.JSON(getErr.Status, getErr)
		return
	}
	
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public")=="true"))
}


func UpdateUser(c *gin.Context) {
	
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {		
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User		
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")		
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	
	isPartial := c.Request.Method == http.MethodPatch
	
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public")=="true"))
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}


func DeleteUser(c *gin.Context) {
	
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {		
		c.JSON(idErr.Status, idErr)
		return
	}
   
	if  err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string {"status":"deleted"})
}

func SearchUsers(c *gin.Context){
	status := c.Query("status")

	users, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	
	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public")=="true"))
}
