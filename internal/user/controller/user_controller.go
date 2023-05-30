package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := userController.userService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
