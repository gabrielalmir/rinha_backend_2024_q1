package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/clientes/:id/extrato", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.JSON(200, gin.H{"id": id})
	})

	println("Server is running on port 8080")
	log.Fatalln(r.Run(":8080"))
}
