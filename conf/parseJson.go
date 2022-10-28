package conf

import (
	"strconv"
	"strings"
)

// возвращает сколько было потрачено газа помесячно
func (g *ModelGas) EveryMonths() (map[int]float64, error) {
	value := make(map[int]float64, 9)
	for _, item := range g.Ethereum.Transactions {
		month, err := takeMonths(item.Time)
		if err != nil {
			return nil, err
		}
		value[month] += item.GasValue
	}
	return value, nil
}

// возвращает число месяца из даты
func takeMonths(fullDate string) (int, error) {
	arr := strings.Split(fullDate, " ")
	dateSlice := strings.Split(arr[0], "-")
	return strconv.Atoi(dateSlice[1])
}

// возвращает среднюю цену покупки газа за день
func (g ModelGas) AveragePricePerDay() map[string]float64 {
	value := make(map[string]float64, 30*8)
	sumGasValuePerDay := make(map[string]float64, 30*8)
	for _, item := range g.Ethereum.Transactions {
		arr := strings.Split(item.Time, " ")
		date := arr[0]
		value[date] += item.GasPrice * item.GasValue
		sumGasValuePerDay[date] += item.GasValue
	}
	for key, _ := range value {
		value[key] /= sumGasValuePerDay[key]
	}
	return value
}

// возвращает сколько было потрачено за весь период
func (g ModelGas) TotalSpent() float64 {
	var value float64
	for _, item := range g.Ethereum.Transactions {
		value += item.GasPrice * item.GasValue
	}
	return value
}

// Частотное распределение цены по часам(за весь период).
func (g ModelGas) TakePriceDistPerHour() (map[int]float64, error) {
	value := make(map[int]float64, 24)
	sumGasValuePerHour := make(map[int]float64, 24)
	for _, item := range g.Ethereum.Transactions {
		date := strings.Split(item.Time, " ")
		hourMin := strings.Split(date[1], ":")
		hour, err := strconv.Atoi(hourMin[0])
		if err != nil {
			return nil, err
		}
		value[hour] += item.GasPrice * item.GasValue
		sumGasValuePerHour[hour] += item.GasValue
	}
	for key, _ := range value {
		value[key] /= sumGasValuePerHour[key]
	}
	return value, nil
}
