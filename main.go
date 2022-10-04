package main

import (
	"log"
	"os"
)

// {
// "time": "22-01-01 00:00", year months day
// "gasPrice": 84.24001654397276,
// "gasValue": 364.7976616468897,
// "average": 0.006750761716697318,
// "maxGasPrice": 15555.000008704,
// "medianGasPrice": 79.535685632
// },

func main() {
	str := "gas_price.json"
	file, err := os.ReadFile(str)
	if err == nil {
		log.Fatalln(err, "file not open")
	}

}

//1) Сколько было потрачено gas помесячно.
//rest api
//web socket
