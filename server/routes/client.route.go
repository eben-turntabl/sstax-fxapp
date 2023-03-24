package routes

import (
	"example/server/controllers"

	"github.com/gin-gonic/gin"
)

type ClientRouteController struct{
	clientController controllers.ClientController
}

func NewClientRouteController(clientController controllers.ClientController) ClientRouteController{
	return ClientRouteController{clientController}
}

func (rc *ClientRouteController) ClientRoute(rg *gin.RouterGroup){
	router := rg.Group("/client")

	router.POST("/exchange", rc.clientController.ConvertBalance)
	router.POST("/purchase", rc.clientController.PlaceOrder)

	router.POST("/withdraw", rc.clientController.WithdrawBalance)

	router.GET("/order/status/:id",rc.clientController.CheckOrder)
}





































