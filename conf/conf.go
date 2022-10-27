package conf

import (
	"encoding/json"
	"io"
	"os"
)

//	type configI interface {
//		LoadConfig(fileName string)
//	}

func New() *ModelgGas {
	return &ModelgGas{}
}

func (a *ModelgGas) LoadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	slyceBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(slyceBytes, a)
	if err != nil {
		return
	}
}
