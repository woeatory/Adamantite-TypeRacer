package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/model/DTO"
	"github.com/woeatory/Adamantite-TypeRacer/internal/user/service"
	"log"
	"net/http"
)

type UserController struct {
	userService service.UserServiceInterface
}

func NewUserController(userService service.UserServiceInterface) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := userController.userService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (userController *UserController) GetUserByID(c *gin.Context) {
	user, err := userController.userService.GetByID(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (userController *UserController) ChangeUsername(c *gin.Context) {
	var input DTO.UserChangeUsernameDto
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting user_id"})
		return
	}
	err := userController.userService.ChangeUsername(input.NewUsername, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session.Set("user_id", input.NewUsername)
	err = session.Save()
	if err != nil {
		log.Printf("error saving new user_id in session storage: %e", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Changed Username"})
}

func (userController *UserController) DeleteUser(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting user_id"})
		return
	}

	err := userController.userService.DeleteUser(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	err = session.Save()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted User"})
}
