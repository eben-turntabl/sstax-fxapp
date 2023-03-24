package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID  primitive.ObjectID `json:"id" bson:"_id"`
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
}

type ConvData struct {
	Email    string `json:"email" bson:"email"`
	FromCurr string `json:"from_curr" bson:"from_curr"`
	ToCurr   string `json:"to_curr" bson:"to_curr"`
	Amount   int    `json:"amount" bson:"amount"`
}

type Stocks struct {
	StockTicker   string `json:"stock_ticker" bson:"stock_ticker"`
	StockName     string `json:"stock_name" bson:"stock_name"`
	StockQuantity int    `json:"stock_quantity" bson:"stock_quantity"`
}
type Porfolio struct {
	PorfolioID string
	Stocks     []Stocks `json:"stocks" bson:"stocks"`
}









func ConvertToWallet(dbResp DBResponse) Wallet {
	return Wallet{
		ID: dbResp.ID,
		USD: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}{
			Amount:   dbResp.Wallet.USD.Amount,
			IsActive: dbResp.Wallet.USD.IsActive,
		},
		GHS: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}{
			Amount:   dbResp.Wallet.GHS.Amount,
			IsActive: dbResp.Wallet.GHS.IsActive,
		},
		KES: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}{
			Amount:   dbResp.Wallet.KES.Amount,
			IsActive: dbResp.Wallet.KES.IsActive,
		},
		NGN: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
		}{
			Amount:   dbResp.Wallet.NGN.Amount,
			IsActive: dbResp.Wallet.NGN.IsActive,
		},
	}
}
