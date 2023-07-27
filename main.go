package main

import (
	"awesomeProject/pkg/PFC"
	"awesomeProject/pkg/PFC/CLI"
	"awesomeProject/pkg/poem"
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
	//err = migrateSchema(config.Conn)
	//if err != nil {
	//	panic(err)
	//}
	//
	//gin.ForceConsoleColor()
	//gin.SetMode(gin.DebugMode)
	//initServices(config)

}

func main() {
	//router := gin.New()
	//initControllers(router)
	//
	//err := router.Run(":8080")
	//if err != nil {
	//	panic(err)
	//}

	//player := PFC.NewPlayer()
	//PFC.NewGame().Play(player.AiPlay())
	histo := PFC.NewHistoric()
	CLI.NewCli().Init(histo)
}

func migrateSchema(db *gorm.DB) error {
	err := db.AutoMigrate(&poem.Poem{})
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
	poem.InitPoemController()

	poem.PoemController(routing)
}

/**
 * initControllers: Init dependency injection needed in the services
 */
func initServices(config Config) {
	poem.InitPoemService(config.Conn)
}
