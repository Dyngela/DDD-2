package main

import (
	"awesomeProject/generic"
	"awesomeProject/mapper"
	"awesomeProject/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Config struct {
	Conn   *gorm.DB
	Logger zerolog.Logger
}

func init() {
	//var err error
	//var config Config
	//
	//config.Conn, err = configuration.ConnectToPostgres()
	//if err != nil {
	//	panic(err)
	//}
	//config.Logger = configuration.InitLogger()
	//config.Logger.Info().Msg("caze")
	//if err = migrateSchema(config.Conn); err != nil {
	//	panic(err)
	//}
	//
	//gin.ForceConsoleColor()
	//gin.SetMode(gin.DebugMode)
	//initServices(config)

}

func main() {
	//mapper.GenerateMapping([]interface{}{user.UserMapper})
	mapper.GenerateMapping()

	//router := gin.New()
	//initControllers(router)
	//docs.SwaggerInfo.BasePath = "/api/v1"
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Run(":8080")
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
func initServices(config Config) {
	user.InitUserService(config.Conn, config.Logger)
	generic.NewGenericRepo(config.Conn)
}
