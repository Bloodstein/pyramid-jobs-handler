package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RunWorker(ctx *gin.Context) {
	queueName := ctx.Param("queue_name")
	if len(queueName) == 0 {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "param 'queue_name' is required",
		})
		return
	}
	h.Services.Launcher.RunWorker(queueName)
}
