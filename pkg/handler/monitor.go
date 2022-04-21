package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Status(ctx *gin.Context) {
	status := h.Services.Monitor.Status()
	ctx.JSON(http.StatusOK, status)
}
