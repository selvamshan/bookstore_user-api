package app

import (
	"github.com/selvamshan/bookstore_user-api/controllers/ping"
	"github.com/selvamshan/bookstore_user-api/controllers/users"
)

func mapUrls (){
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("users/:user_id", users.GetUser)
	//router.GET("/users/search", controllers.SearchUser)

}