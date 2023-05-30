package user

import "github.com/gin-gonic/gin"

var userService IUserService

func InitUserController() {
	userService = NewUserService()
}

func UserController(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/", userService.GetById)
	}
}
