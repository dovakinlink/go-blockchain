package main

import (
	"math"
	"math/big"
)

func main() {
	//c := client.Get()
	//account := common.HexToAddress("")
	//balance, err := c.BalanceAt(context.Background(), account, nil)
	//if err != nil {
	//	panic(err)
	//}
	//
	//ethValue := ConvertToEth(balance)
	//fmt.Println(ethValue)
	//
	//pendingBalance, err := c.PendingBalanceAt(context.Background(), account)
	//pendingEthValue := ConvertToEth(pendingBalance)
	//fmt.Println(pendingEthValue)

}

func ConvertToEth(balance *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	return ethValue
}
