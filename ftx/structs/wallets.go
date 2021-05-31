package structs

type Balances struct {
	Success bool `json:"success"`
	Result []struct {
		Coin string `json:"coin"`
		Free float64 `json:"free"`
		SpotBorrow float64 `json:"spotBorrow"`
		Total float64 `json:"total"`
		USDValue float64 `json:"usdValue"`
		AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
	}
}