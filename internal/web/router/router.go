package router

import (
	"log"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/web/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	RegisterRoutes(router)
	log.Fatalln(router.Run(":8080"))
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", handler.HandleCustomerStatement)
}
