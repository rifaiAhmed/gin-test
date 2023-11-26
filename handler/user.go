package handler

import (
	"net/http"
	"test-gin/helper"
	"test-gin/user"

	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService user.Service
}

func Newuserhandler(userService user.Service) *userhandler {
	return &userhandler{userService}
}

func (h *userhandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	response := helper.APIResponse("Account has been registered!", http.StatusOK, "Succesfully", newUser)
	c.JSON(http.StatusOK, response)
}
