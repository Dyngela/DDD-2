package poem

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
)

var poemService IPoemService

func InitPoemController() {
	poemService = NewPoemService()
}

func PoemController(router *gin.Engine) {
	v1 := router.Group("api/v1/poem")
	{
		v1.GET("/:id", getById)
		v1.PUT("/", create)
	}
}

func getById(ctx *gin.Context) {
	poem, err := poemService.GetById(ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, GetPoemByIdResponse{
		Author:      poem.Author,
		ReleaseDate: poem.ReleaseDate,
		Content:     poem.Content,
	})
}

func create(ctx *gin.Context) {
	err := poemService.Create(ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, common.GenericResponse{Message: "Votre poème a été créé"})
}
