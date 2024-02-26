package router

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/handler"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/logger"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(log *logger.Logger) {
	log.Infof("Setting up the router ...")
	router := gin.Default()
	router.Use(CORSMiddleware())
	RegisterRoutes(router)

	PORT := "8080"
	log.Infof("Starting the server on port :%s", PORT)
	router.Run(":" + PORT)
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", handler.HandleCustomerStatement)
	r.POST("/clientes/:id/transacoes", handler.HandleCreateTrasaction)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
