package model

import "time"

type Data struct {
	Code       string `json:"code" db:"code"`
	Codein     string `json:"codein" db:"codein"`
	Name       string `json:"name" db:"name"`
	High       string `json:"high" db:"high"`
	Low        string `json:"low" db:"low"`
	VarBid     string `json:"varBid" db:"varBid"`
	PctChange  string `json:"pctChange" db:"pctChange"`
	Bid        string `json:"bid" db:"bid"`
	Ask        string `json:"ask" db:"ask"`
	Timestamp  string `json:"timestamp" db:"timestamp"`
	CreateDate string `json:"create_date" db:"create_date"`
}

type ExchangeRate struct {
	Dollar Data `json:"USDBRL"`
	Euro   Data `json:"EURBRL"`
}

type ExchangeRateHistory struct {
	ID         int       `db:"id"`
	Bid        string    `db:"bid"`
	Conversion string    `db:"conversion"`
	CreatedAt  time.Time `db:"created_at"`
}
