package httpserv

import (
	"todos/handler"
	"todos/infra"
	"todos/ports/createtodo"
	"todos/ports/gettodo"
	"todos/ports/gettodos"
	"todos/ports/updatetodo"
	"todos/services/create"
	"todos/services/list"
	"todos/services/update"

	"github.com/gin-gonic/gin"
)

func bindCreateRoute(app *gin.Engine) {
	createTodo := createtodo.NewAdaptorFile(infra.NewFileManager(infra.TodosStorage))
	service := create.New(createTodo)
	hdl := handler.NewCraeteHandler(service)

	app.POST("/v1/create", hdl.Handle)
}

func bindUpdateRoute(app *gin.Engine) {
	getTodo := gettodo.NewAdaptorFile(infra.NewFileManager(infra.TodosStorage))
	updateTodo := updatetodo.NewAdaptorFile(infra.NewFileManager(infra.TodosStorage))
	service := update.New(getTodo, updateTodo)
	hdl := handler.NewUpdateHandler(service)

	app.POST("/v1/update", hdl.Handle)
}

func bindListRoute(app *gin.Engine) {
	getTodos := gettodos.NewAdaptorFile(infra.NewFileManager(infra.TodosStorage))
	service := list.New(getTodos)
	hdl := handler.NewListHandler(service)

	app.GET("/v1/list", hdl.Handle)
}
