package customer

import "github.com/gin-gonic/gin"

var customerService *CustomerService

func HandleCustomerStatement(ctx *gin.Context) {
	id := ctx.Param("id")

	customer, err := customerService.GetCustomerStatement(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, customer)
}
