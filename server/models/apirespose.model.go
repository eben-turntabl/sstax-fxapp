package models

type FXRateResponse struct {
	Result             string `json:"result"`
	Documentation      string `json:"documentation"`
	TermsOfUse         string `json:"terms_of_use"`
	TimeLastUpdateUnix int    `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string `json:"time_last_update_utc"`
	TimeNextUpdateUnix int    `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string `json:"time_next_update_utc"`
	BaseCode           string `json:"base_code"`
	TargetCode         string `json:"target_code"`
	ConversionRate     float32    `json:"conversion_rate"`
	ConversionResult   float32    `json:"conversion_result"`
}

type ExchangeRate struct {
	Rates map[string]int `json:"rates"`
	Base  string         `json:"base"`
}

type ExchangeRequest struct {
	Email    string `json:"email"`
	FromCurr string `json:"from_curr"`
	ToCurr   string `json:"to_curr"`
	Amount   int    `json:"amount"`
}

type ExchangeResponse struct {
	ExchangedAmount int `json:"exchanged_amount"`
}
