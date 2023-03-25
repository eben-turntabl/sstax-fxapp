package controllers

import (
	"example/server/models"
	"example/server/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientService   services.ClientService
	exchangeService services.ExchangeService
	orderService    services.OrderService
}

func NewClientController(clientService services.ClientService, exchangeService services.ExchangeService,orderService services.OrderService) ClientController {
	return ClientController{clientService, exchangeService,orderService}
}

func (cc ClientController) PlaceOrder(ctx *gin.Context) {
	var orderInfo *models.FXOrder

	if err := ctx.ShouldBindJSON(&orderInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
	}
	resp,err := cc.orderService.PurchaseOrder(orderInfo)
	if err != nil {
		fmt.Printf("Order Invalid: %v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})

		return
	} else {
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": "success", "data": resp.ID})

	}




}


func (cc ClientController) CheckOrder(ctx *gin.Context) {
	order_id := ctx.Param("id")
	

	resp,err := cc.orderService.CheckOrder(order_id)
	if err != nil {
		fmt.Printf("Order Number is  Invalid: %v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})

		return
	} else {
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": "success", "data": resp})

	}


}
func (cc ClientController) WithdrawBalance(ctx *gin.Context) {
		var withdrawalInfo	*models.CashOut

		if err := ctx.ShouldBindJSON(&withdrawalInfo); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		}

		resp, err := cc.exchangeService.WithdrawBalance(withdrawalInfo)
	if err != nil {
		fmt.Printf("error withdrawing balance: %v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})

		return
	} else {
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": "success", "Message":"Withdrawal to account ending "+withdrawalInfo.BankDetails.Account.BankName+" successful","data":resp })

	}


}

func (cc ClientController) ConvertBalance(ctx *gin.Context) {
	var mconvData *models.ConvData

	if err := ctx.ShouldBindJSON(&mconvData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
	}

	if cc.exchangeService == nil {
		fmt.Println("exchangeService is nil")

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "exchangeService is nil"})

		return
	}

	resp, err := cc.exchangeService.ConvertBalance(mconvData)
	if err != nil {
		fmt.Printf("error converting balance: %v\n", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})

		return
	} else {
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": "success", "data": resp})

	}

}
