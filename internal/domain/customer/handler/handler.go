package handler

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/domain/customer/service"
	"github.com/gin-gonic/gin"
)

var customerService *service.CustomerService

func HandleCustomerStatement(ctx *gin.Context) {
	id := ctx.Param("id")

	customer, err := customerService.GetCustomerStatement(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, customer)
}
