package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734"
var ganacheURL = "127.0.0.1:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a ether a client:%v", err)
	}
	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	fmt.Println("Block Number: ", block.Number())

	addr := "0x95222290DD7278Aa3Ddd389Cc1E1d165CC4BAfe5"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to ger a Balance:%v", err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	BalacneEth := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("Balance: ", BalacneEth)

}
