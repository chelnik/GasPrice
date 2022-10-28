package conf

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
