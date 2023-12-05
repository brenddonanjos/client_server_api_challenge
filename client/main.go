package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8080/cotacao", nil)
	if err != nil {
		log.Println(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Request time limit exceeded. (client)")
		}
		log.Println(err.Error())
	}
	defer res.Body.Close()

	// read json response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading body response:", err)
	}
	log.Println(string(body))
}
