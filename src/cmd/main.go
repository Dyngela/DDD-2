package main

import (
	"awesomeProject/src/configuration"
	"awesomeProject/src/pkg/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	var err error
	var db *gorm.DB

	db, err = configuration.ConnectToPostgres()
	if err != nil {
		panic(err)
	}
	if err = migrateSchema(db); err != nil {
		panic(err)
	}

	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)
	initServices(db)
}

func main() {
	service := user.NewUserService()
	service.GetById(nil)
	router := gin.New()
	initControllers(router)
}

func migrateSchema(db *gorm.DB) error {
	err := db.AutoMigrate(&user.Users{}, &user.Roles{})
	if err != nil {
		return err
	}
	return nil
}

/*
*   initControllers: Init dependency injection needed in the controllers
*    Then and only then we can initialize the API endpoints
 */
func initControllers(routing *gin.Engine) {
	user.InitUserController()

	user.UserController(routing)
}

/**
 * initControllers: Init dependency injection needed in the services
 */
func initServices(db *gorm.DB) {
	user.InitUserService(db)
}
