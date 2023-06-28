package user

import (
	"github.com/gin-gonic/gin"
)

var userService IUserService

func InitUserController() {
	userService = NewUserService()
}

func UserController(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/:id", getById)
	}
}

// @Summary Get a user by its ID
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} Users
// @Failure 400 {object} common.APIError
// @Router /example/helloworld [get]
func getById(context *gin.Context) {
	userService.GetById()
}
