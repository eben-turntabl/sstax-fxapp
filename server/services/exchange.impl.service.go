package services

import (
	"context"
	"errors"
	"example/server/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type ExchangeServiceImpl struct{
	collection *mongo.Collection
	ctx	context.Context
	
}

func NewExchangeService(collection *mongo.Collection, ctx context.Context) ExchangeService{
	return &ExchangeServiceImpl{collection, ctx }
}

func (ec *ExchangeServiceImpl) ConvertBalance(exchange *models.ConvData)(*models.Wallet,error){

reqUri := "https://v6.exchangerate-api.com/v6/38a115464aa968dd3f9af5d0/pair/" + exchange.FromCurr +"/"+exchange.ToCurr+"/"+fmt.Sprintf("%v", exchange.Amount)
resp, err := http.Get(reqUri)
if err != nil {
	return nil, errors.New("Unable to make requests")
}
fmt.Println(resp)
return nil,nil

}

func (ec *ExchangeServiceImpl) UpdateBalance(exchange *models.ConvData)(*models.Wallet,error){
	
}

