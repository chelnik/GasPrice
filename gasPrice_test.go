package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	bytes := workWithData()
	data := &Response{
		Monthly:          nil,
		PricePerDay:      nil,
		PriceDistPerHour: nil,
		EntirePeriod:     0,
	}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		t.Error("Unmarshal error")
	}
	url := "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json"
	_, err = http.Get(url)
	if err != nil {
		t.Error("Get error")
	}
}

func TestGetQuery(t *testing.T) {
	url := "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json"
	_, err := http.Get(url)
	if err != nil {
		t.Error("Get error")
	}
}
