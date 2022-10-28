package conf

import (
	"encoding/json"
	"io"
	"os"
)

//	type configI interface {
//		LoadConfig(fileName string)
//	}

func New() *ModelGas {
	return &ModelGas{}
}

func (g *ModelGas) LoadConfig(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	slyceBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(slyceBytes, g)
	if err != nil {
		return
	}
}
