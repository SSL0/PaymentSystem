// @title PaymentAPI
// @version 1.0
// @description API for wallet transactions and balance management
// @BasePath /api

// @tag.name Transactions
// @tag.description "Endpoints for transaction operations"

// @tag.name Wallet
// @tag.description "Endpoints for wallet information"

package handler

import (
	"PaymentAPI/pkg/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		api.POST("/send", h.Send)
		api.GET("/transactions", h.GetLast)
		api.GET("/wallet/:address/balance", h.GetBalance)

	}

	return router
}
