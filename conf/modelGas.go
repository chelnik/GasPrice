package conf

import (
	"log"
	"strconv"
	"strings"
)

type ModelGas struct {
	Ethereum struct {
		Transactions []struct {
			Time           string  `json:"time"`
			GasPrice       float64 `json:"gasPrice"`
			GasValue       float64 `json:"gasValue"`
			Average        float64 `json:"average"`
			MaxGasPrice    float64 `json:"maxGasPrice"`
			MedianGasPrice float64 `json:"medianGasPrice"`
		} `json:"transactions"`
	} `json:"ethereum"`
}

func (g *ModelGas) EveryMonths() map[int]float64 {
	value := make(map[int]float64, 9)
	for _, item := range g.Ethereum.Transactions {
		month, err := takeMonths(item.Time)
		if err != nil {
			log.Println("err in take months", err)
		}
		value[month] += item.GasValue
	}
	return value
}

func takeMonths(fullDate string) (int, error) {
	arr := strings.Split(fullDate, " ")
	dateSlice := strings.Split(arr[0], "-")
	return strconv.Atoi(dateSlice[1])
}

func (g ModelGas) AveragePricePerDay() map[string]float64 {
	value := make(map[string]float64, 30*8)
	// days := make(map[string]float64, 30*8)
	sumGasValuePerDay := make(map[string]float64, 30*8)
	for _, item := range g.Ethereum.Transactions {
		arr := strings.Split(item.Time, " ")
		date := arr[0]
		value[date] += item.GasPrice * item.GasValue
		sumGasValuePerDay[date] += item.GasValue
		// days[date]++
	}
	for key, _ := range value {
		value[key] /= sumGasValuePerDay[key]
	}
	return value
}

// fullDate := "22-01-01 00:00"
// arr := strings.Split(fullDate, " ")
// // date := arr[0]
// hour := strings.Split(arr[1], ":")
// fmt.Println(hour)
// newbie := strings.Split(date, "-")
