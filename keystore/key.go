package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	/*var password = "password"
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	pw, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pw.Address)*/

	password := "password"
	b, err := ioutil.ReadFile("./keystore/wallet/UTC--2024-07-12T16-59-20.831170000Z--9108c80a07d129db728aa8de2203220aa874a0c4")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private: ", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public: ", hexutil.Encode(pData))

	fmt.Println("ADD: ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
