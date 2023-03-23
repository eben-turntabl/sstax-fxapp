package models 


type Wallet struct {
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
		

}

type ConvData struct{
	FromCurr	string	`json:"from_curr" bson:"from_curr"`
	ToCurr		string	`json:"to_curr" bson:"to_curr"`
	Amount		float32	`json:"amount" bson:"amount"`
}