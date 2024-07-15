package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"

	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://sepolia.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734"


func main() {
	/*ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := ks.NewAccount("password")
	if err != nil {
		log.Fatal(err)
	}
	_, err = ks.NewAccount("password")
	if err != nil {
		log.Fatal(err)
	}

	"1b1cdb6a828cfebaf6b16c9ce2b062d362e023ac"
	"624041a2b96f407ae04cca88b5655cb472932395"*/

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	addr1 := common.HexToAddress("1b1cdb6a828cfebaf6b16c9ce2b062d362e023ac")
	addr2 := common.HexToAddress("624041a2b96f407ae04cca88b5655cb472932395")
	b1, err := client.BalanceAt(context.Background(), addr1, nil)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := client.BalanceAt(context.Background(), addr2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance:", b1)
	fmt.Println("Balance:", b2)
	nonce, err := client.PendingNonceAt(context.Background(), addr1)
	if err != nil{
		log.Fatal(err)
	}

	amount := big.NewInt(10000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, addr2, amount, 21000, gasPrice, nil)
	ChainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(ChainID)
	}

	br, err := ioutil.ReadFile("./wallet/UTC--2024-07-12T18-05-07.686775300Z--624041a2b96f407ae04cca88b5655cb472932395")
	if err != nil {
		log.Fatal(err)
	}

	keyP, err := keystore.DecryptKey(br, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err = types.SignTx(tx, types.NewEIP2930Signer(ChainID), keyP.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("send tx: %s", tx.Hash().Hex())

}
