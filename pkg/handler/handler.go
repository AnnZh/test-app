package handler

import (
	"github.com/AnnZh/test-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/get-data", h.getData)

		queries := api.Group("/queries")
		{
			queries.GET("/over-speed", h.getOverspeedCars)
			queries.GET("/min-max", h.getMinMaxSpeedCars)
		}
	}

	return router
}
