package library

import (
	"awesomeProject/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

var libService ILibService

func InitLibController() {
	libService = NewLibService()
}

func LibController(router *gin.Engine) {
	v1 := router.Group("api/v1/book")
	{
		v1.POST("/borrow/:id", borrow)
		v1.GET("/", getAll)
		v1.GET("/:id", getById)
		v1.DELETE("/:id", deleteById)
		v1.PUT("/", create)
	}
}

func create(context *gin.Context) {
	var book Book
	_ = context.ShouldBindJSON(&book)
	err := libService.Create(book, context)
	if err != nil {
		context.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	context.JSON(200, common.GenericResponse{Message: "Book created"})

}

func deleteById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := libService.Delete(id, ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, common.GenericResponse{Message: "Book deleted"})

}

func getById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := libService.FindById(id, ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, book)
}

func getAll(ctx *gin.Context) {
	books, err := libService.FindAll(ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, books)
}

func borrow(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := libService.Borrow(id, ctx)
	if err != nil {
		ctx.JSON(500, common.APIError{Message: err.Error()})
		return
	}
	ctx.JSON(200, common.GenericResponse{Message: "Book has been checked has borrowed"})
}
