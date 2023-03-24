package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientLogin struct {
	Email    string `json:"email" bson:"email"  binding:"required`
	Password string `json:"password" bson:"password" binding:"required`
}

type ClientSignup struct {
	FirstName string    `json:"first_name" bson:"first_name" binding:"required`
	LastName  string    `json:"last_name" bson:"last_name" binding:"required`
	Email     string    `json:"email" bson:"email"  binding:"required`
	Role      string    `json:"role" bson:"role" `
	Password  string    `json:"password" bson:"password"  binding:"required`
	CreatedAt time.Time `json:"created_at" bson:"created_at" `
	Wallet    struct {
		USD struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"USD" bson:"USD"`
		GHS struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"GHS" bson:"GHS"`
		KES struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"KES" bson:"KES"`
		NGN struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"NGN" bson:"NGN"`
	} `json:"wallet" bson:"wallet"`

	BankDetails struct {
		Account1 struct {
			BankName  string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`
		} `json:"account1" bson:"account1"`
		Account2 struct {
			BankName  string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`
		} `json:"account2" bson:"account2"`
	} `json:"bank_details" bson:"bank_details"`

	Orders     []string `json:"orders,omitempty" bson:"orders,omitempty" default:"[]" `
	PorfolioID string   `json:"portfolio_id" bson:"portfolio_id"`
}


type UserUpdate struct {
	FirstName string    `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Role      string    `json:"role,omitempty" bson:"role,omitempty"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Wallet    struct {
		USD struct {
			Amount   int  `json:"amount,omitempty" bson:"amount,omitempty"`
			IsActive bool `json:"is_active,omitempty" bson:"is_active,omitempty"`
		} `json:"USD,omitempty" bson:"USD,omitempty"`
		GHS struct {
			Amount   int  `json:"amount,omitempty" bson:"amount,omitempty"`
			IsActive bool `json:"is_active,omitempty" bson:"is_active,omitempty"`
		} `json:"GHS,omitempty" bson:"GHS,omitempty"`
		KES struct {
			Amount   int  `json:"amount,omitempty" bson:"amount,omitempty"`
			IsActive bool `json:"is_active,omitempty" bson:"is_active,omitempty"`
		} `json:"KES,omitempty" bson:"KES,omitempty"`
		NGN struct {
			Amount   int  `json:"amount,omitempty" bson:"amount,omitempty"`
			IsActive bool `json:"is_active,omitempty" bson:"is_active,omitempty"`
		} `json:"NGN,omitempty" bson:"NGN,omitempty"`
	} `json:"wallet,omitempty" bson:"wallet,omitempty"`

	BankDetails struct {
		Account1 struct {
			BankName  string `json:"bank_name,omitempty" bson:"bank_name,omitempty"`
			AcctNumbr string `json:"acctnumbr,omitempty" bson:"acctnumbr,omitempty"`
		} `json:"account1,omitempty" bson:"account1,omitempty"`
		Account2 struct {
			BankName  string `json:"bank_name,omitempty" bson:"bank_name,omitempty"`
			AcctNumbr string `json:"acctnumbr,omitempty" bson:"acctnumbr,omitempty"`
		} `json:"account2,omitempty" bson:"account2,omitempty"`
	} `json:"bank_details" bson:"bank_details,omitempty"`

	Orders     []string `json:"orders,omitempty" bson:"orders,omitempty" default:"[]"`
	PorfolioID string   `json:"portfolio_id,omitempty" bson:"portfolio_id,omitempty"`
}

//Remember to add constraints

type DBResponse struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" binding:"required`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Email     string             `json:"email" bson:"email"`
	Role      string             `json:"role" bson:"role"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Wallet    struct {
		USD struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"USD" bson:"USD"`
		GHS struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"GHS" bson:"GHS"`
		KES struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"KES" bson:"KES"`
		NGN struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		} `json:"NGN" bson:"NGN"`
	} `json:"wallet" bson:"wallet"`

	BankDetails struct {
		Account1 struct {
			BankName  string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`
		} `json:"account1" bson:"account1"`
		Account2 struct {
			BankName  string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`
		} `json:"account2" bson:"account2"`
	} `json:"bank_details" bson:"bank_details"`

	Orders     []string `json:"orders,omitempty" bson:"orders,omitempty" default:"[]"`
	PorfolioID string   `json:"portfolio_id" bson:"portfolio_id"`
}

// {"first_name":"ben","last_name":"john","email":"kskskdj@jk.com","password"::"tonga"}
