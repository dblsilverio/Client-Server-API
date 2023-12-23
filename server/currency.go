package main

type CurrencyResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type CurrentyTO struct {
	Bid float32 `json:"bid"`
}
