package services

import "example/server/models"


type ExchangeService interface{
	

	ConvertBalance(*models.ConvData) (*models.DBResponse,error)
	WithdrawBalance(*models.CashOut)	(*models.DBResponse,error)
}
