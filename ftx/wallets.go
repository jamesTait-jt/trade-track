package ftx

import (
	"log"
)
import (
	"github.com/jamesTait-jt/trade-track/ftx/structs"
)

type Balances structs.Balances

func (client *FtxClient) GetBalances() (Balances, error) {
	var balances Balances
	resp, err := client._get("/wallet/balances", []byte(""))
	if err != nil {
		log.Printf("Error GetCoins", err)
		return balances, err
	}
	err = client._processResponse(resp, &balances)
	return balances, err
}