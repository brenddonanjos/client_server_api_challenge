package main

import (
	"net/http"

	"github.com/brenddonanjos/client_server_api_challenge/server/handler"
)

func main() {
	http.HandleFunc("/cotacao", handler.GetDollarRate)
	http.HandleFunc("/history", handler.GetExchangeRateHistory)
	http.ListenAndServe(":8080", nil)
}
