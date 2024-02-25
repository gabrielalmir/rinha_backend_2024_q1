package main

import (
	"log"

	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.SetupRoutes(r)

	println("Server is running on port 8080")
	log.Fatalln(r.Run(":8080"))
}
