package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StopWorker(ctx *gin.Context) {
	guid := ctx.Param("worker_guid")

	if len(guid) == 0 {
		log.Println("Worker's guid not sent")
	}

	h.Services.Aborter.StopWorker(guid)
}
