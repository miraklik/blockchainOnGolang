package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pkey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pdata := crypto.FromECDSA(pkey)
	fmt.Println(hexutil.Encode(pdata))

	puData := crypto.FromECDSAPub(&pkey.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	fmt.Println(crypto.PubkeyToAddress(pkey.PublicKey).Hex())
}
