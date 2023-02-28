package model

type MoneyChanger struct {
	Money_changer_ID string  `json:"money_changer_ID"`
	Name             string  `json:"name"`
	Exchange_rate    float64 `json:"exchange_rate"`
	Currency_type    string  `json:"currency_type"`
	Country          string  `json:"country"`
}
