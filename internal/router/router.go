package router

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()
	RegisterRoutes(router)
	log.Fatalln(router.Run(":8080"))
}
