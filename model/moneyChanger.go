package model

type MoneyChanger struct {
	Money_Changer_ID string `json:"money_Changer_ID"`
	Name string `json:"name"`
	Exchange_Rate float64 `json:"exchange_Rate"`
	Currency_Type string `json:"currency_Type"`
	Country string `json:"country"`
}

