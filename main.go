package main

import (
	"encoding/json"
	"fmt"

	"github.com/chelnik/GasPriceTask/conf"
)

// 1) Сколько было потрачено gas помесячно.
const fileName = "gas_price.json"

type Response struct {
	Monthly          map[int]float64    // monthly expenses
	PricePerDay      map[string]float64 // average price per day
	PriceDistPerHour int                // price distribution per hour
	EntirePeriod     int                // payment for the entire period
}

func main() {
	data := conf.New()
	data.LoadConfig(fileName)
	response := Response{
		Monthly:          data.EveryMonths(),
		PricePerDay:      data.AveragePricePerDay(),
		PriceDistPerHour: 0,
		EntirePeriod:     0,
	}
	_, err := json.Marshal(response)
	if err != nil {
		return
	}
	// fmt.Println(response)

}
func testPricePerDay(response *Response) {
	for i, i2 := range response.PricePerDay {
		if "22-08-05" == i {
			fmt.Println(i, i2)
		}
	}
}
