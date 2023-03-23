package models

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type ClientLogin struct{
	Email		string `json:"email" bson:"email"  binding:"required`
	Password	string	`json:"password" bson:"password" binding:"required`



}



type ClientSignup struct{
	FirstName	string		`json:"first_name" bson:"first_name" binding:"required`
	LastName	string		`json:"last_name" bson:"last_name" binding:"required`
	Email		string		`json:"email" bson:"email"  binding:"required`
	Role		string		`json:"role" bson:"role" `
	Password	string		`json:"password" bson:"password"  binding:"required`
	CreatedAt	time.Time	`json:"created_at" bson:"created_at" `
	Wallet		struct {
		USD struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"USD" bson:"USD"`
		GHS	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"GHS" bson:"GHS"`
		KES	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"KES" bson:"KES"`
		NGN	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"NGN" bson:"NGN"`
		
	}`json:"wallet" bson:"wallet"`

	BankDetails	struct{
		Account1 struct{
			BankName string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`

		}`json:"account1" bson:"account1"`
		Account2 struct{
			BankName string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`

			
		}`json:"account2" bson:"account2"`


	}`json:"bank_details"`
}

//Remember to add constraints


type DBResponse struct{
	ID	primitive.ObjectID	`json:"id" bson:"_id" binding:"required`
	FirstName	string		`json:"first_name" bson:"first_name"`
	LastName	string		`json:"last_name" bson:"last_name"`
	Email		string		`json:"email" bson:"email"`
	Role		string		`json:"role" bson:"role"`
	Password	string		`json:"password" bson:"password"`
	CreatedAt	time.Time	`json:"created_at" bson:"created_at"`
	Wallet		struct {
		USD struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"USD" bson:"USD"`
		GHS	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"GHS" bson:"GHS"`
		KES	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"KES" bson:"KES"`
		NGN	struct{
			Amount int `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}`json:"NGN" bson:"NGN"`
		
	}`json:"wallet" bson:"wallet"`

	BankDetails	struct{
		Account1 struct{
			BankName string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`

		}`json:"account1" bson:"account1"`
		Account2 struct{
			BankName string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`

			
		}`json:"account2" bson:"account2"`


	}`json:"bank_details"`
}



// {"first_name":"ben","last_name":"john","email":"kskskdj@jk.com","password"::"tonga"}