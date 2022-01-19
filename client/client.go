package client

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io/")
	if err != nil {
		panic(err)
	}
	fmt.Println("rpc connected")
}

func Get() *ethclient.Client {
	return client
}
