package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello, world!"})
	})

	println("Server is running on port 8080")
	log.Fatalln(r.Run(":8080"))
}
