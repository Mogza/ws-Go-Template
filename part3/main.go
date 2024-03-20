package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

var AlchemySepoliaURL = "https://eth-sepolia.g.alchemy.com/v2/A-a8SZHrZcgIGjS05NhDAumHdLLrSXEq"

func main() {
	client, err := ethclient.DialContext(context.Background(), AlchemySepoliaURL)
	if err != nil {
		log.Fatalf("Error while creating a client : %v\n", err)
	}
	defer client.Close()

}

func printBalance(client *ethclient.Client) {
	// Etape 2 : Ajoutez votre code ici
}

func sendTx(client *ethclient.Client, toAddress common.Address, fromAddress common.Address, amount *big.Int) {
	// Etape 3 : Ajoutez votre code ici
}
