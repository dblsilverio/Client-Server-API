package main

import (
	"fmt"
	"os"
)

const serverUrl = "http://localhost:8080/cotacao"

func main() {

	to := RequestCurrencyData()
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
