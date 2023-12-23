package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func RequestCurrencyData() CurrencyTO {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", serverUrl, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)

	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var to CurrencyTO

	err = json.Unmarshal(data, &to)
	if err != nil {
		panic(err)
	}
	return to
}
