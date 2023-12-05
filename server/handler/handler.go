package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/brenddonanjos/client_server_api_challenge/server/model"
	"github.com/brenddonanjos/client_server_api_challenge/server/service"
)

func GetDollarRate(w http.ResponseWriter, r *http.Request) {

	//get exchange rate from external api
	exchangeRate, err := service.GetExchangeRateFromAwesomeApi()
	if err != nil {
		fmt.Println("Error getting exchange rate :", err.Error())
		http.Error(w, "Error getting exchange rate:  "+err.Error(), http.StatusBadRequest)
		return
	}

	//store on sqlite
	history := &model.ExchangeRateHistory{
		Bid:        exchangeRate.Dollar.Bid,
		Conversion: exchangeRate.Dollar.Name,
		CreatedAt:  time.Now(),
	}
	err = service.StoreExchangeRateHistory(*history)
	if err != nil {
		fmt.Println("Error storing exchange rate on DB:", err.Error())
		http.Error(w, "Error storing exchange rate on DB: "+err.Error(), http.StatusBadRequest)
		return
	}

	//write on file
	file, err := os.Create("./files/dollar_exchange.txt")
	if err != nil {
		fmt.Println("Error creating txt File:", err.Error())
		http.Error(w, "Error creating txt File: "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = file.WriteString("Dólar: " + exchangeRate.Dollar.Bid)
	if err != nil {
		fmt.Println("Error writing on txt File:", err.Error())
		http.Error(w, "Error writing on txt File  ", http.StatusInternalServerError)
		return
	}
	file.Close()

	//returns json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(exchangeRate.Dollar.Bid)
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Error encoding JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetExchangeRateHistory(w http.ResponseWriter, r *http.Request) {
	exchangesRates, err := service.GetAllExchangeRateHistory()
	if err != nil {
		fmt.Println("Error getting exchanges rates :", err.Error())
		http.Error(w, "Error getting exchanges rates ", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(exchangesRates)
	if err != nil {
		fmt.Println("Error encoding JSON response:", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}
