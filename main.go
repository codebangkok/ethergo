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
	client, err := ethclient.Dial("https://kovan.infura.io/v3/24595eef73ac4734b3dd3eea707a6f3c")
	if err != nil {
		log.Fatal(err)
	}

	myAccount := common.HexToAddress("0x008f9b59Eb93b9762f304992B94F65fCe0571776")
	// checkEthBalance(*client, myAccount)
	_ = myAccount

	contractAddress := common.HexToAddress("0x899F7782AB089E7206769d028Ce51894dbA40813")
	bondCoinContract, err := bondcoin.NewBondcoin(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// checkBalance(*client, *bondCoinContract, myAccount)

	privateKey, err := crypto.HexToECDSA("fa4613e7cca618a14acba25d131f466e68d3330010a639f5528635b6f4a7fb80")
	if err != nil {
		log.Fatal(err)
	}

	recipientAccount := common.HexToAddress("0x64c80aC601507c7864D9E31215b2E66824172C7b")
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
