package main

import (
	"test-gin/db"
	"test-gin/handler"
	"test-gin/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.CreateCon()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userhandler := handler.Newuserhandler(userService)
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/create", userhandler.RegisterUser)
	router.Run(":9191")
}
