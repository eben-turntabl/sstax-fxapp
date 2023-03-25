package services

import (
	"context"
	"example/server/models"
	"fmt"

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


 
func (os *OrderServiceImpl) PurchaseOrder(orderInfo *models.FXOrder)(*models.FXOrderDBResponse,error){

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
	var newOrder *models.FXOrderDBResponse
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
	return newOrder,nil
	

}

func (os *OrderServiceImpl) PlaceOrder(orderInfo *models.ClientOrder) (*models.OrderDBResponse, error) {

	
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

func (os *OrderServiceImpl) CheckOrder(orderId string) (*models.FXOrderDBResponse, error) {


	var orderStatus *models.FXOrderDBResponse
	oid, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}
	err = os.collection.FindOne(os.ctx, filter).Decode(&orderStatus)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.FXOrderDBResponse{}, err
		}
		return nil, err
	}
	return orderStatus, nil
}
