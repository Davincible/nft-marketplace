package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func transferToAll(conn *ethclient.Client, from accounts.Account, amount int) error {
	for _, acc := range kstore.Accounts() {
		if transfer && acc.Address.String() != coinbaseAddr {
			log.Printf("Sending transaction of %d BNB, for shits 'n giggles, to %s", amount, acc.Address.Hex())

			_, err := sendTransaction(conn, from, acc, transferAmount)
			if err != nil {
				log.Printf("Sending transaction failed: %v", err)
				continue
			}

			time.Sleep(3 * time.Second)
		}

		if err := logBalance(conn, acc.Address); err != nil {
			return err
		}
	}

	return nil
}

func sendTransaction(client *ethclient.Client, from, to accounts.Account, amount int) (*common.Hash, error) {
	value := new(big.Int).Mul(big.NewInt(int64(amount)), big.NewInt(params.Ether))
	gasLimit := uint64(21000)

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return nil, err
	}

	feeCap, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, feeCap, nil)

	signedTx, err := kstore.SignTx(from, tx, chainid)
	if err != nil {
		return nil, fmt.Errorf("sign transaction: %w", err)
	}

	if err := client.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, fmt.Errorf("send transaction: %w", err)
	}

	txHash := tx.Hash()

	return &txHash, nil
}
