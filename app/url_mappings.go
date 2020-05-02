package app

import (
	"github.com/selvamshan/bookstore_user-api/controllers/ping"
	"github.com/selvamshan/bookstore_user-api/controllers/users"
)

func mapUrls (){
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.GET("/internal/users/search", users.SearchUsers)
	router.POST("/users/login", users.Login)
}