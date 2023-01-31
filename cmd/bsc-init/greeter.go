package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	_, tx, greeter, err := contracts.DeployGreeter(auth, conn, "Hello World!")
	if err != nil {
		return fmt.Errorf("deploy greeter: %w", err)
	}

	if _, err = waitForReceipt(conn, tx.Hash()); err != nil {
		return err
	}

	greeting, err := greeter.Greet(nil)
	if err != nil {
		return fmt.Errorf("call greeter: %w", err)
	}

	log.Printf("Received greeting from smart contract: %s!", greeting)

	return nil
}

func waitForReceipt(conn *ethclient.Client, addr common.Hash) (*types.Receipt, error) {
	deadline := time.After(time.Second * 180)

	for {
		select {
		case <-deadline:
			return nil, errors.New("failed to fetch tx receipt")
		default:
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
			defer cancel()

			r, err := conn.TransactionReceipt(ctx, addr)
			if err != nil && errors.Is(err, ethereum.NotFound) {
				time.Sleep(time.Second)
				continue
			} else if err != nil {
				return nil, fmt.Errorf("fetch receipt: %w", err)
			}

			return r, nil
		}
	}
}
