package main

import (
	"awesomeProject/src/configuration"
	"awesomeProject/src/pkg/user"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	var err error
	err = configuration.ConnectToPostgres()
	if err != nil {
		panic(err)
	}
	if err = migrateSchema(); err != nil {
		panic(err)
	}

	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

func migrateSchema() error {
	err := configuration.GetConfig().Conn.AutoMigrate(&user.Users{}, &user.Roles{})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	log.Println("Ok")
	user.Test()
	router := gin.New()
	initControllers(router)
}

func initControllers(routing *gin.Engine) {
	user.UserController(routing)
}
