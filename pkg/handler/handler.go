package handler

import (
	"github.com/Bloodstein/pyramid-jobs-handler/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Services service.Service
}

func NewHandler(s service.Service) Handler {
	return Handler{
		Services: s,
	}
}

func (h *Handler) Routes() *gin.Engine {

	router := gin.New()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("help", h.Help)
			v1.GET("status", h.Status)
			workers := v1.Group("workers")
			{
				launch := workers.Group("launch")
				{
					launch.GET(":queue_name", h.RunWorker)
				}
				stop := workers.Group("stop")
				{
					stop.GET(":worker_guid", h.StopWorker)
				}
			}
			jobs := v1.Group("jobs")
			{
				dispatch := jobs.Group("dispatch")
				{
					dispatch.GET(":queue_name", h.RunWorker)
				}
				monitoring := jobs.Group("report")
				{
					monitoring.GET(":jobs_state", h.RunWorker)
				}
			}
		}
	}

	return router
}
