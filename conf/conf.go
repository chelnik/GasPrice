package conf

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func New() *ModelGas {
	return &ModelGas{}
}

// подгружает gas_price.json из интернета
func (g *ModelGas) GetRequest() error {
	url := "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json"
	getGasPriceJson, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http.Get error %s", err)
	}
	sliceBytes, err := io.ReadAll(getGasPriceJson.Body)
	err = json.Unmarshal(sliceBytes, g)
	if err != nil {
		return fmt.Errorf("Unmarshal error %s", err)
	}
	return nil
}

// подгружает gas_price.json из файла
func (g *ModelGas) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open error %s", err)
	}
	sliceBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("readAll error %s", err)
	}
	err = json.Unmarshal(sliceBytes, g)
	if err != nil {
		return fmt.Errorf("unmarshal error %s", err)
	}
	return nil
}
