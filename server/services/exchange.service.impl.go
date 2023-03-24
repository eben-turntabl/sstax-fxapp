package services

import (
	"context"
	"encoding/json"

	"example/server/models"
	"example/server/utils"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ExchangeServiceImpl struct{
	collection *mongo.Collection
	ctx	context.Context
}

func NewExchangeService(collection *mongo.Collection, ctx context.Context) ExchangeService{
	return &ExchangeServiceImpl{collection, ctx }
}


func (es *ExchangeServiceImpl)	ConvertBalance(convData *models.ConvData) (*models.DBResponse,error){
	var walletData *models.DBResponse 
	filter := bson.M{"email":strings.ToLower(convData.Email)}
	err := es.collection.FindOne(es.ctx,filter).Decode(&walletData)
	if err != nil {
		if err == mongo.ErrNoDocuments{
			return &models.DBResponse{},err
		}
		return nil,err

	}

	currentBal := models.ConvertToWallet(*walletData)

	reqUri := "https://v6.exchangerate-api.com/v6/38a115464aa968dd3f9af5d0/pair/" + convData.FromCurr +"/"+convData.ToCurr+"/"+fmt.Sprintf("%v", convData.Amount)
 // Change the base currency if necessary
	resp, err := http.Get(reqUri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var exchangeRate models.FXRateResponse
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRate); err != nil {
		return nil, err
	}
	

	switch convData.ToCurr {
	case "GHS":
		currentBal.GHS.Amount = currentBal.GHS.Amount  + int(exchangeRate.ConversionResult)
			
	case "USD":

		currentBal.USD.Amount = currentBal.USD.Amount  + int(exchangeRate.ConversionResult)

	case "KES":
		currentBal.KES.Amount = currentBal.KES.Amount  + int(exchangeRate.ConversionResult)

	case "NGN":
		currentBal.NGN.Amount = currentBal.NGN.Amount  + int(exchangeRate.ConversionResult)

	}


	switch convData.FromCurr {
	case "GHS":
		currentBal.GHS.Amount = currentBal.GHS.Amount - (convData.Amount)

	case "USD":
		currentBal.USD.Amount = currentBal.USD.Amount - (convData.Amount)
	
	case "KES":
		currentBal.KES.Amount = currentBal.KES.Amount - (convData.Amount)

	case "NGN":
		currentBal.NGN.Amount = currentBal.NGN.Amount - (convData.Amount)
	}
	data,err := utils.ToDoc(currentBal)
	if err != nil{
		return &models.DBResponse{},err
	}
	updated:= bson.D{{Key: "$set",Value: bson.D{{Key: "wallet", Value: data}}}}
	query:= bson.D{{Key: "_id",Value: walletData.ID}}

	updateData := es.collection.FindOneAndUpdate(es.ctx,query,updated,options.FindOneAndUpdate().SetReturnDocument(1))
	
	var result *models.DBResponse
	result = &models.DBResponse{} // initialize result pointer
	if err := updateData.Decode(result); err != nil{
		
		return nil, err
	}
		
	return result, nil
	


	

	// return walletData,nil
}


func (es *ExchangeServiceImpl) WithdrawBalance(withdrawalInfo *models.CashOut) (*models.DBResponse, error) {
	var userInfo *models.DBResponse
	oid, _ := primitive.ObjectIDFromHex(withdrawalInfo.ClientID)
	filter := bson.M{"_id": oid}

	err := es.collection.FindOne(es.ctx, filter).Decode(&userInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	userInfo.Wallet.GHS.Amount = userInfo.Wallet.GHS.Amount - withdrawalInfo.Amount

	data, err := utils.ToDoc(userInfo)
	if err != nil {
		return nil, err
	}
	updated := bson.D{{Key: "$set", Value: data}}
	query := bson.D{{Key: "_id", Value: userInfo.ID}}

	// Use options.After to return the updated document
	updateData := es.collection.FindOneAndUpdate(es.ctx, query, updated, options.FindOneAndUpdate().SetReturnDocument(options.After))

	result := &models.DBResponse{} // initialize result pointer
	if err := updateData.Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
