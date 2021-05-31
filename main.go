package main

import (
	"fmt"
	"os"
)
import (
	"github.com/jamesTait-jt/trade-track/ftx"
)

// -- GLOBAL VARIABLES
// 1. FirstUpperCase = public
// 2. firstLowerCase = package private


func main() {
	client := ftx.New(os.Getenv("FTX_API_KEY"), os.Getenv("FTX_API_SECRET"), "")
	// client.GetTrades()
	fmt.Println(client.GetBalances())
}