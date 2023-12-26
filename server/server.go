package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	exchangeMux := http.NewServeMux()
	exchangeMux.HandleFunc("/cotacao", ExchangeHandler)

	err := http.ListenAndServe(":8080", exchangeMux)
	if err != nil {
		panic(fmt.Sprintf("Failed server boot: %s", err))
	}

}

func ExchangeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	exchange, err := FetchExchange()
	if err != nil {
		log.Printf("Exchange request failed: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	entity, err := SaveExchangeInfo(exchange)
	if err != nil {
		log.Printf("SaveExchange Info failed: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(CurrentyTO{Bid: entity.Bid})
	if err != nil {
		fmt.Printf("Encoding response failed: %s", err)
	}
}
