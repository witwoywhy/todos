package handler

import (
	"net/http"
	"todos/services/list"

	"github.com/gin-gonic/gin"
)

type listHandler struct {
	service list.Service
}

func NewListHandler(service list.Service) *listHandler {
	return &listHandler{
		service: service,
	}
}

func (h *listHandler) Handle(ctx *gin.Context) {
	var request list.Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	response, err := h.service.Execute(request)
	if err != nil {
		ctx.JSON(err.GetHttpStatus(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}
