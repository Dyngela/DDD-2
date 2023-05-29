package user

import (
	"awesomeProject/src/configuration"
	"github.com/gin-gonic/gin"
	"log"
)

func userRepo() *UserRepo {
	return NewUserRepo(configuration.GetConfig().Conn)
}

func Test() {
	user, err := userRepo().GetAllUser()
	if err != nil {
		panic(err)
	}
	log.Println(user)
}

func getAll(context *gin.Context) {

}
