package handler

import (
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/dto"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/customer/service"
	"github.com/gabrielalmir/rinha_backend_2024_q1/internal/db"
	"github.com/gin-gonic/gin"
)

func HandleCustomerStatement(ctx *gin.Context) {
	id := ctx.Param("id")

	dbConn := db.Instance
	customerService := service.NewCustomerService(dbConn)

	customer, err := customerService.GetCustomerStatement(id)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.Status(404)
			return
		}

		ctx.Status(500)
	}

	ctx.JSON(200, customer)
}

func HandleCreateTrasaction(ctx *gin.Context) {
	var transactionDTO dto.TransactionDTO
	id := ctx.Param("id")

	dbConn := db.Instance
	customerService := service.NewCustomerService(dbConn)

	if err := ctx.ShouldBindJSON(&transactionDTO); err != nil {
		ctx.Status(400)
		return
	}

	customer, err := customerService.CreateTransaction(id, transactionDTO)

	if err != nil {
		if err.Error() == "record not found" {
			ctx.Status(404)
			return
		}

		if err.Error() == "invalid transaction" {
			ctx.Status(422)
			return
		}

		ctx.Status(500)
	}

	ctx.JSON(200, customer)
}
