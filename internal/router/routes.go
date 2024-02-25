package router

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	r.GET("/clientes/:id/extrato", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{"id": id})
	})
}
