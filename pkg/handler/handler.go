package handler

import (
	"PaymentAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/send", h.Send)
		api.GET("/transactions", h.GetLast)
		api.GET("/wallet/:address/balance", h.GetBalance)
	}

	return router
}
