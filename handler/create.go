package handler

import (
	"net/http"
	"todos/services/create"

	"github.com/gin-gonic/gin"
)

type createHandler struct {
	service create.Service
}

func NewCraeteHandler(service create.Service) *createHandler {
	return &createHandler{
		service: service,
	}
}

func (h *createHandler) Handle(ctx *gin.Context) {
	var request create.Request
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
