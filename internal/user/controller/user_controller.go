package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model/DTO"
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

func (userController *UserController) GetUserByID(c *gin.Context) {
	user, err := userController.userService.GetByID(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (userController *UserController) ChangeUsername(c *gin.Context) {
	var input DTO.ChangeUsernameDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	var userID string
	v := session.Get("id")
	if v == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	} else {
		userID = v.(string)
	}
	err := userController.userService.ChangeUsername(input.NewUsername, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully changed username"})
}

func (userController *UserController) DeleteUser(c *gin.Context) {

	session := sessions.Default(c)
	var userID string
	v := session.Get("id")
	if v == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unauthorized"})
		return
	} else {
		userID = v.(string)
	}
	err := userController.userService.DeleteUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session.Delete("id")
	err = session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
