package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func weiToEth(w *big.Int) string {
	return new(big.Rat).SetFrac(w, big.NewInt(1e18)).FloatString(18)
}

func getEnv(env, str string) string {
	if len(os.Getenv(env)) > 0 {
		return os.Getenv(env)
	}

	return str
}

func logBalance(conn *ethclient.Client, addr common.Address) error {
	balance, err := conn.PendingBalanceAt(context.Background(), addr)
	if err != nil {
		return fmt.Errorf("fetch balance: %w", err)
	}

	log.Printf("Balance of %s: %s BNB", addr.Hex(), weiToEth(balance))

	return nil
}
