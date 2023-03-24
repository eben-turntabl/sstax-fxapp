package services

import (
	"context"
	"errors"
	"example/server/models"
	"example/server/utils"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientServiceImpl struct{
collection *mongo.Collection
ctx context.Context
}

func NewClientService(collection *mongo.Collection,ctx context.Context) ClientService{
	return &ClientServiceImpl{collection, ctx}
}

func(cs *ClientServiceImpl) FindUserByEmail(email string)	(*models.DBResponse, error){
	var user *models.DBResponse
	query := bson.M{"email":strings.ToLower(email)}
	err := cs.collection.FindOne(cs.ctx,query).Decode(&user)

	if err != nil{
		if err == mongo.ErrNoDocuments{
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}
func (cs *ClientServiceImpl) FindUserById(id string) (*models.DBResponse, error){

	oid, _ := primitive.ObjectIDFromHex(id)
	var user *models.DBResponse
	query:= bson.M{"_id":oid}
	err := cs.collection.FindOne(cs.ctx,query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments{
			return &models.DBResponse{},err
		}
		return nil, err

	}
	return user, nil

}

func (cs *ClientServiceImpl) UpdateUserInfoEmail(email string ,update *models.UserUpdate) (*models.DBResponse, error){

	data,err := utils.ToDoc(update)
	if err != nil{
		return &models.DBResponse{},err
	}


	filter := bson.D{{Key: "email",Value: email}}

	change := bson.D{{Key: "$set",Value: data}}
	result := cs.collection.FindOneAndUpdate(cs.ctx, filter, change,options.FindOneAndUpdate().SetReturnDocument(1))
	
	var returnDocument *models.DBResponse
	if err := result.Decode(&returnDocument); err != nil {
		return nil,errors.New("User doesnt exist")
	}
return returnDocument,nil 
}


func (cs *ClientServiceImpl) UpdateUserInfo(id string ,update *models.UserUpdate) (*models.DBResponse, error){

	data,err := utils.ToDoc(update)
	if err != nil{
		return &models.DBResponse{},err
	}
	oid ,_ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id",Value: oid}}

	change := bson.D{{Key: "$set",Value: data}}
	result := cs.collection.FindOneAndUpdate(cs.ctx, filter, change,options.FindOneAndUpdate().SetReturnDocument(1))
	
	var returnDocument *models.DBResponse
	if err := result.Decode(&returnDocument); err != nil {
		return nil,errors.New("User doesnt exist")
	}
return returnDocument,nil 
}
