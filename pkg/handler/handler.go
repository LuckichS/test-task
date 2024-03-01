package handler

import (
	"Hezzl/pkg/service"

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

	good := router.Group("/good")
	{
		good.POST("/create", h.create)
		good.PATCH("/update", h.update)
		good.DELETE("/remove", h.remove)
		good.GET("/list", h.list)
		good.PATCH("/reprioritiize", h.reprioritiize)
	}

	return router
}
