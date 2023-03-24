package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientOrder struct {
	Symbol     string    `json:"symbol" bson:"symbol"`
	PlacedAt   time.Time `json:"placed_at" bson:"placed_at"`
	Price      int       `json:"price" bson:"price"`
	Side       string    `json:"side" bson:"side"`
	Ticker     string    `json:"ticker" bson:"ticker"`
	Status     string    `json:"status" bson:"status"`
	Quantity   int       `json:"quantity" bson:"quantity"`
	PorfolioID string    `json:"portfolio_id" bson:"portfolio_id"`
	ClientID   string    `json:"client_id" bson:"client_id"`
}

type OrderDBResponse struct {
	ID         primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Symbol     string             `json:"symbol" bson:"symbol"`
	PlacedAt   time.Time          `json:"placed_at" bson:"placed_at"`
	Price      int                `json:"price" bson:"price"`
	Side       string             `json:"side" bson:"side"`
	Ticker     string             `json:"ticker" bson:"ticker"`
	Status     string             `json:"status" bson:"status"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	PorfolioID string             `json:"portfolio_id" bson:"portfolio_id"`
	ClientID   string             `json:"client_id" bson:"client_id"`
}

type OrderUpdate struct {
	Symbol     string    `json:"symbol,omitempty" bson:"symbol,omitempty"`
	PlacedAt   time.Time `json:"placed_at,omitempty" bson:"placed_at,omitempty"`
	Price      int       `json:"price,omitempty" bson:"price,omitempty"`
	Side       string    `json:"side,omitempty" bson:"side,omitempty"`
	Ticker     string    `json:"ticker,omitempty" bson:"ticker,omitempty"`
	Status     string    `json:"status,omitempty" bson:"status,omitempty"`
	Quantity   int       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	PorfolioID string    `json:"portfolio_id,omitempty" bson:"portfolio_id,omitempty"`
	ClientID   string    `json:"client_id" bson:"client_id"`
}

type CashOut struct {
	Amount      int `json:"amount" bson:"amount"`
	BankDetails struct {
		Account struct {
			BankName  string `json:"bank_name" bson:"bank_name"`
			AcctNumbr string `json:"acctnumbr" bson:"acctnumbr"`
		} `json:"account1" bson:"account"`
	} `json:"bank_details" bson:"bank_details"`
	ClientID string `json:"client_id" bson:"client_id"`
}


