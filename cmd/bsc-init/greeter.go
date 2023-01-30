package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Davincible/nft-marketplace/contracts"
)

func getGreeting(conn *ethclient.Client, from accounts.Account) error {
	log.Println("Deploying Greeter contract")

	id, err := conn.ChainID(context.Background())
	if err != nil {
		return fmt.Errorf("fetch chainID: %w", err)
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(kstore, from, id)
	if err != nil {
		log.Fatal(err)
	}

	_, _, greeter, err := contracts.DeployGreeter(auth, conn, "Hello World!")
	if err != nil {
		return fmt.Errorf("deploy greeter: %w", err)
	}

	time.Sleep(time.Second * 5)

	greeting, err := greeter.Greet(nil)
	if err != nil {
		return fmt.Errorf("call greeter: %w", err)
	}

	log.Printf("Received greeting from smart contract: %s!", greeting)

	return nil
}
