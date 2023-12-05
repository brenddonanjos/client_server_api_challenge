package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/brenddonanjos/client_server_api_challenge/server/database"
	"github.com/brenddonanjos/client_server_api_challenge/server/model"
)

func StoreExchangeRateHistory(history model.ExchangeRateHistory) error {
	db, err := database.StartConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO exchange_rate_history(conversion, bid, created_at) VALUES ($1, $2, $3)")
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return errors.New("store db time limit exceeded (Database)")
		}
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, history.Bid, history.Conversion, history.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func GetAllExchangeRateHistory() (histories []model.ExchangeRateHistory, err error) {
	db, err := database.StartConnection()
	if err != nil {
		return
	}
	defer db.Close()
	resultados, err := db.Query("select * from exchange_rate_history order by id desc")
	if err != nil {
		log.Fatal(err)
	}
	defer resultados.Close()

	for resultados.Next() {
		var history model.ExchangeRateHistory
		err := resultados.Scan(&history.ID, &history.Conversion, &history.Bid, &history.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		histories = append(histories, history)
	}
	if err = resultados.Err(); err != nil {
		return
	}

	return histories, nil
}

func GetExchangeRateFromAwesomeApi() (exchangeRate model.ExchangeRate, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return exchangeRate, errors.New("request time limit exceeded (awesomeapi)")
		}
		return
	}
	defer res.Body.Close()

	// read json response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading body response:", err)
	}
	//Prepare and return struct
	err = json.Unmarshal(body, &exchangeRate)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return exchangeRate, nil
}
