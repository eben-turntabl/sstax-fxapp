package services

import "example/server/models"


type ExchangeService interface{
	
	
	ConvertBalance(*models.ConvData) (*models.Wallet,error)
}
