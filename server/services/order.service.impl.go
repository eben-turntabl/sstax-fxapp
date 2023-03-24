package services

import (
	"context"
	"errors"
	"example/server/models"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
	collection2	*mongo.Collection
}

func NewOrderService(collection *mongo.Collection, ctx context.Context, collection2	*mongo.Collection) OrderService {
	return &OrderServiceImpl{collection, ctx,collection2}
}

func (os *OrderServiceImpl) PlaceOrder(orderInfo *models.ClientOrder) (*models.OrderDBResponse, error) {

	reqUri := "https://v6.exchangerate-api.com/v6/38a115464aa968dd3f9af5d0/pair/USD/GHS"
	resp, err := http.Get(reqUri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("exchange rate API returned non-200 status code")
	}
	orderInfo.Status= "Pending"
	orderInfo.PlacedAt=time.Now()

	res, err := os.collection.InsertOne(os.ctx, orderInfo)
	fmt.Println(res)
	fmt.Println(err)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 1100 {
			return nil, er
		}
		return nil, err
	}

	fmt.Println(err)
	var newOrder *models.OrderDBResponse
	query := bson.M{"_id": res.InsertedID}

	err = os.collection.FindOne(os.ctx, query).Decode(&newOrder)
	if err != nil {
		return nil, err
	}

	updated:= bson.D{{Key: "$push",Value: bson.D{{Key: "orders", Value: newOrder.ID.String()}}}}
	clientID, err := primitive.ObjectIDFromHex(orderInfo.ClientID)
	if err != nil {
		fmt.Println("Error")
	}
	filter := bson.M{"_id": clientID}

	updateData := os.collection2.FindOneAndUpdate(os.ctx,filter,updated,options.FindOneAndUpdate().SetReturnDocument(1))
	
	var result *models.DBResponse
	result = &models.DBResponse{} // initialize result pointer
	if err := updateData.Decode(result); err != nil{
		
		return nil, err
	}
	return newOrder, nil
}

func (os *OrderServiceImpl) CheckOrder(orderId string) (*models.OrderDBResponse, error) {
	var orderStatus *models.OrderDBResponse
	oid, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	err = os.collection.FindOne(os.ctx, filter).Decode(&orderStatus)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.OrderDBResponse{}, err
		}
		return nil, err
	}
	return orderStatus, nil
}
