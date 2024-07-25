package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/longtk26/go-ecommerce/internal/services"
	"github.com/longtk26/go-ecommerce/pkg/response"
)

type UserController struct {
	userService *services.UserService
}


func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 2001, uc.userService.GetUserByID())
}