package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const currencyUrl = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

func FetchExchange() (CurrencyResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", currencyUrl, nil)
	if err != nil {
		return CurrencyResponse{}, err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(fmt.Sprintf("Closing Body reader failed: %s", err))
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CurrencyResponse{}, err
	}

	var currency CurrencyResponse
	err = json.Unmarshal(body, &currency)
	if err != nil {
		return CurrencyResponse{}, err
	}

	return currency, nil
}
