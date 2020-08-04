package main

import (
	"context"
	"crypto/ecdsa"
	"ethergo/bondcoin"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	myAccount := common.HexToAddress("")
	// checkEthBalance(*client, myAccount)
	_ = myAccount

	contractAddress := common.HexToAddress("")
	bondCoinContract, err := bondcoin.NewBondcoin(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// checkBalance(*client, *bondCoinContract, myAccount)

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	recipientAccount := common.HexToAddress("")
	transfer(*client, *bondCoinContract, *privateKey, recipientAccount, *big.NewInt(100))
}

func transfer(client ethclient.Client, bondCoinContract bondcoin.Bondcoin, privateKey ecdsa.PrivateKey, recipient common.Address, amount big.Int) {
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	options := bind.NewKeyedTransactor(&privateKey)
	options.Value = big.NewInt(0)
	options.GasPrice = gasPrice
	options.GasLimit = 300000

	txn, err := bondCoinContract.Transfer(options, recipient, &amount)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction: https://kovan.etherscan.io/tx/%s\n", txn.Hash().Hex())
}

func checkBalance(client ethclient.Client, bondCoinContract bondcoin.Bondcoin, account common.Address) {
	options := bind.CallOpts{From: account}

	balance, err := bondCoinContract.BalanceOf(&options, account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("BON Balance:", balance)
}

func checkEthBalance(client ethclient.Client, account common.Address) {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ETH Balance:", balance)
}
