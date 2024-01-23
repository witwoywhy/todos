package httpserv

import (
	"todos/handler"
	"todos/infra"
	"todos/ports/createtodo"
	"todos/services/create"

	"github.com/gin-gonic/gin"
)

func bindCreateRoute(app *gin.Engine) {
	createTodo := createtodo.NewAdaptorFile(infra.NewFileManager(infra.TodosStorage))
	service := create.New(createTodo)
	hdl := handler.NewCraeteHandler(service)

	app.POST("/v1/create", hdl.Handle)
}
