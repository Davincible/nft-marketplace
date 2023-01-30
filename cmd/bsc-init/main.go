package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	addr         = getEnv("ADDRESS", "/tmp/geth.ipc")
	keystoreDir  = getEnv("KEYSTORE_DIR", "./accounts")
	coinbaseAddr = getEnv("COINBASE_ADDR", "0x7fd60C817837dCFEFCa6D0A52A44980d12F70C59")

	transfer       = true
	transferAmount = 2

	kstore = keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
)

func main() {
	log.Printf("Starting init container, connecting to '%s'...", addr)
	conn := connect(addr)
	defer conn.Close()

	log.Printf("Accounts loaded: %d\n", len(kstore.Accounts()))

	coinbase := accounts.Account{Address: common.HexToAddress(coinbaseAddr)}

	if err := kstore.Unlock(coinbase, ""); err != nil {
		log.Fatalf("Failed to unlock coinbase account")
	}

	for _, acc := range kstore.Accounts() {
		if transfer && acc.Address.String() != coinbaseAddr {
			log.Printf("Sending transaction of %d BNB, for shits 'n giggles, to %s", transferAmount, acc.Address.Hex())

			_, err := sendTransaction(conn, coinbase, acc, transferAmount)
			if err != nil {
				log.Printf("Sending transaction failed: %v", err)
				continue
			}

			time.Sleep(3 * time.Second)
		}

		balance, err := conn.PendingBalanceAt(context.Background(), acc.Address)
		if err != nil {
			log.Fatalf("Failed to fetch balance: %v", err)
		}

		log.Printf("Balance of %s: %s BNB", acc.Address.Hex(), weiToEth(balance))
	}
}

func connect(addr string) *ethclient.Client {
	for {
		client, err := ethclient.Dial(addr)
		if err != nil {
			log.Println("Waiting for BSC node to be ready...")

			time.Sleep(4 * time.Second)

			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		id, err := client.NetworkID(ctx)
		if err != nil {
			log.Println(err)
			client.Close()
			time.Sleep(4 * time.Second)

			continue
		}

		log.Printf("Connected to Network '%x'", id.Bytes())

		return client
	}
}

func getEnv(env, str string) string {
	if len(os.Getenv(env)) > 0 {
		return os.Getenv(env)
	}

	return str
}

func weiToEth(w *big.Int) string {
	return new(big.Rat).SetFrac(w, big.NewInt(1e18)).FloatString(18)
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
