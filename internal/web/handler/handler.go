package handler

import "github.com/gin-gonic/gin"

func HandleCustomerStatement(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(200, gin.H{"id": id})
}
