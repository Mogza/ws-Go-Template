package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var AlchemyURL = "https://eth-mainnet.g.alchemy.com/v2/1GmTcH4hWwRvXbD9AZsomxovwFjawYn8"

func main() {
	client, err := ethclient.DialContext(context.Background(), AlchemyURL)
	if err != nil {
		log.Fatalf("Error while creating a client : %v\n", err)
	}
	defer client.Close()

	// Ajoutez votre code ici
}
