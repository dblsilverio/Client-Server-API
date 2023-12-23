package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const serverUrl = "http://localhost:8080/cotacao"

func main() {

	to := requestCurrencyData()
	saveCurrencyData(to)

}

func saveCurrencyData(to CurrencyTO) {
	cotacaoFile, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(cotacaoFile)

	_, err = cotacaoFile.WriteString(fmt.Sprintf("Dólar: %f", to.Bid))
	if err != nil {
		panic(err)
	}

}

func requestCurrencyData() CurrencyTO {
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
