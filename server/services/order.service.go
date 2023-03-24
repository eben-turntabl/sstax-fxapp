package services

import "example/server/models"


type OrderService interface{

	PlaceOrder(*models.ClientOrder)(*models.OrderDBResponse,error)
	CheckOrder(string)(*models.OrderDBResponse,error)

}