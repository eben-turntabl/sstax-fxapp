package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wallet struct {
	ID  primitive.ObjectID `json:"id" bson:"_id"`
	USD struct {
		Amount   int  `json:"amount" bson:"amount"`
		IsActive bool `json:"is_active" bson:"is_active"`
		Currency string `json:"currency" bson:"currency" default:"USD"`
	} `json:"USD" bson:"USD"`
	GHS struct {
		Amount   int  `json:"amount" bson:"amount"`
		IsActive bool `json:"is_active" bson:"is_active"`
		
		Currency string `json:"currency" bson:"currency" default:"GHS"`
	} `json:"GHS" bson:"GHS"`
	KES struct {
		Amount   int  `json:"amount" bson:"amount"`
		IsActive bool `json:"is_active" bson:"is_active"`
		
		Currency string `json:"currency" bson:"currency" default:"KES"`
	} `json:"KES" bson:"KES"`
	NGN struct {
		Amount   int  `json:"amount" bson:"amount"`
		IsActive bool `json:"is_active" bson:"is_active"`
		Currency string `json:"currency" bson:"currency" default:"NGN"`
	} `json:"NGN" bson:"NGN"`
}


type Currency struct {
	Symbol string `json:"symbol" bson:"symbol"`
	Amount float64 `json:"amount" bson:"amount"`

}  



type ConvData struct {
	Email    string `json:"email" bson:"email"`
	FromCurr string `json:"from_curr" bson:"from_curr" default:"GHS"`
	ToCurr   string `json:"to_curr" bson:"to_curr"`
	Amount   int    `json:"amount" bson:"amount"`
}








func ConvertToWallet(dbResp DBResponse) Wallet {
	return Wallet{
		ID: dbResp.ID,
		USD: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
			Currency string `json:"currency" bson:"currency" default:"USD"`
		}{
			Amount:   dbResp.Wallet.USD.Amount,
			IsActive: dbResp.Wallet.USD.IsActive,
			Currency: dbResp.Wallet.USD.Currency,
			
		},
		GHS: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
			Currency string `json:"currency" bson:"currency" default:"GHS"`
		}{
			Amount:   dbResp.Wallet.GHS.Amount,
			IsActive: dbResp.Wallet.GHS.IsActive,
			Currency: dbResp.Wallet.GHS.Currency,

			
		},
		KES: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
			Currency string `json:"currency" bson:"currency" default:"KES"`
		}{
			Amount:   dbResp.Wallet.KES.Amount,
			IsActive: dbResp.Wallet.KES.IsActive,
			Currency: dbResp.Wallet.KES.Currency,
			
		},
		NGN: struct {
			Amount   int  `json:"amount" bson:"amount"`
			IsActive bool `json:"is_active" bson:"is_active"`
			Currency string `json:"currency" bson:"currency" default:"NGN"`
		}{
			Amount:   dbResp.Wallet.NGN.Amount,
			IsActive: dbResp.Wallet.NGN.IsActive,
			Currency: dbResp.Wallet.NGN.Currency,

		},
	}
}
