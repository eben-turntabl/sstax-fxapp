package services

import (
	"context"
	"errors"
	"example/server/models"
	"example/server/utils"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type AuthServiceImpl struct{
	collection *mongo.Collection
	ctx	context.Context
	
}

func NewAuthService(collection *mongo.Collection, ctx context.Context) AuthService{
	return &AuthServiceImpl{collection, ctx }
}

func (uc *AuthServiceImpl) UserRegister(user *models.ClientSignup) (*models.DBResponse, error){
	user.CreatedAt=time.Now()
	user.Email=strings.ToLower(user.Email)
	user.Role="user"
	user.Wallet.GHS.Currency="GHS"
	user.Wallet.NGN.Currency="NGN"
	user.Wallet.KES.Currency="KES"
	user.Wallet.USD.Currency="USD"
	user.Wallet.GHS.IsActive = true
	user.Wallet.GHS.Amount= 50000
	user.Wallet.USD.Amount= 200000
	user.Wallet.USD.IsActive = true


	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	fmt.Println(user)


	
	_ , err := uc.FindUserByEmail(user.Email)
	
	
	// if err != nil {
	// 	return nil, errors.New("Email already exists")
	// }

	res, err := uc.collection.InsertOne(uc.ctx,&user)
	fmt.Println(res)
	fmt.Println(err)
	if err !=nil{
		if er, ok:= err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 1100{
			return nil, errors.New("Email already exists")
		}
		return nil,err
	}


	

	

	fmt.Println(res)
	fmt.Println(err)
	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}


	err = uc.collection.FindOne(uc.ctx,query).Decode(&newUser)
	if err !=nil {
		return nil,err
	}
	return newUser,nil
}

func (uc *AuthServiceImpl) UserLogin(user *models.ClientLogin , )(*models.DBResponse, error){
	
	fmt.Println(user)
	
	login , err := uc.FindUserByEmail(user.Email)
	
	
	if err != nil {

		return nil, errors.New("Email not found")
		
	}
	userInfo := utils.ValidatePassword(login.Password,user.Password)
	if userInfo != nil{

		return nil, errors.New("Invalid Credentials")

	}
	return login, nil
}







func(uc *AuthServiceImpl) FindUserByEmail(email string)	(*models.DBResponse, error){
	var user *models.DBResponse
	query := bson.M{"email":strings.ToLower(email)}
	
	err := uc.collection.FindOne(uc.ctx,query).Decode(&user)

	if err != nil{
		if err == mongo.ErrNoDocuments{
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}
func (uc *AuthServiceImpl) FindUserById(id string) (*models.DBResponse, error){

	oid, _ := primitive.ObjectIDFromHex(id)
	var user *models.DBResponse
	query:= bson.M{"_id":oid}
	err := uc.collection.FindOne(uc.ctx,query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments{
			return &models.DBResponse{},err
		}
		return nil, err

	}
	return user, nil

}