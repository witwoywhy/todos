package handler

import (
	"net/http"
	"todos/services/update"

	"github.com/gin-gonic/gin"
)

type updateHandler struct {
	service update.Service
}

func NewUpdateHandler(service update.Service) *updateHandler {
	return &updateHandler{
		service: service,
	}
}

func (h *updateHandler) Handle(ctx *gin.Context) {
	var request update.Request
	if err := ctx.BindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	err := h.service.Execute(request)
	if err != nil {
		ctx.JSON(err.GetHttpStatus(), err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
