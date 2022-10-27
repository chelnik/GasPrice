package main

import (
	"encoding/json"
	"log"
	"os"
)

// 1) Сколько было потрачено gas помесячно.

func main() {
	mapa := make(map[int]ModelgGas)
	str := "gas_price.json"
data:
	file, err := os.ReadFile(str)
	if err == nil {
		log.Fatalln(err, "file not open")
	}
	err := json.Unmarshal(file, &mapa)
	if err != nil {
		return
	}

}
