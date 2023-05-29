package user

import "github.com/gin-gonic/gin"

func UserController(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/", getAll)
	}
}
