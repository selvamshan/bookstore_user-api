package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"	
	"strconv"	
	"github.com/selvamshan/bookstore_user-api/domain/users"
	"github.com/selvamshan/bookstore_user-api/services"
	"github.com/selvamshan/bookstore_utils-go/rest_errors"
	"github.com/selvamshan/bookstore_oauth-go/oauth"
)

func getUserId(userIdParam string)(int64, *rest_errors.RestErr) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		return 0, rest_errors.NewBadRequestError("user id should be number")		
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User	
	
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")
		
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
    if err := oauth.Authenticate(c.Request); err != nil {
		fmt.Println(err)
		c.JSON(err.Status, err)
		return
	}
	
	if callerId := oauth.GetCallerId(c.Request); callerId == 0 {
		err := rest_errors.RestErr{
			Status: http.StatusUnauthorized,
			Message: "resource not available",
		}
		c.JSON(err.Status, err.Message)
		return
	}

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
	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshal(false))
		return
	}
	
	c.JSON(http.StatusOK, user.Marshal(oauth.IsPublic(c.Request)))
}


func UpdateUser(c *gin.Context) {
	
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {		
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User		
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := rest_errors.NewBadRequestError("Invalid json body")		
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

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr:= rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, err := services.UsersService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public")=="true"))

}