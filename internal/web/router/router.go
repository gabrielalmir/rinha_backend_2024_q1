package web

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	RegisterRoutes(router)
	log.Fatalln(router.Run(":8080"))
}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{"id": id})
	})
}
