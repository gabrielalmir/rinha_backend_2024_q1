package router

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/customer/handler"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/logger"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(log *logger.Logger) {
	log.Infof("Setting up the router ...")
	router := gin.Default()
	RegisterRoutes(router)

	PORT := "8080"
	log.Infof("Starting the server on port :%s", PORT)
	router.Run(":" + PORT)
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", handler.HandleCustomerStatement)
}
