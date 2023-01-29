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

	transfer = false

	k = keystore.NewKeyStore(keystoreDir, keystore.StandardScryptN, keystore.StandardScryptP)
)

func main() {
	log.Printf("Starting init container, connecting to '%s'...", addr)
	conn := connect(addr)
	defer conn.Close()

	log.Printf("Accounts loaded: %d\n", len(k.Accounts()))

	coinbase := accounts.Account{Address: common.HexToAddress(coinbaseAddr)}

	if err := k.Unlock(coinbase, ""); err != nil {
		log.Fatalf("Failed to unlock coinbase account")
	}

	for _, acc := range k.Accounts() {
		if transfer && acc.Address.String() != coinbaseAddr {
			tx, err := sendTransaction(conn, coinbase, acc, 2)
			if err != nil {
				log.Printf("Send transaction from %s to %s: %v", coinbase.Address.Hex(), acc.Address.Hex(), err)
				continue
			}

			if err := waitTransaction(conn, *tx); err != nil {
				log.Fatalf("Wait for transaction: %v", err)
			}

			time.Sleep(time.Second)
		}

		balance, err := conn.BalanceAt(context.Background(), acc.Address, nil)
		if err != nil {
			log.Fatalf("Failed to fetch balance: %v", err)
		}

		log.Printf("Balance of %s: %s ETH", acc.Address.Hex(), weiToEth(balance))
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

func sendTransaction(cl *ethclient.Client, from, to accounts.Account, amount int) (*common.Hash, error) {
	value := new(big.Int).Mul(big.NewInt(int64(amount)), big.NewInt(params.Ether))
	gasLimit := uint64(21000)

	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	nonce, err := cl.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return nil, err
	}

	tipCap, err := cl.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	feeCap, err := cl.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	tx := types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   chainid,
			Nonce:     nonce,
			GasTipCap: tipCap,
			GasFeeCap: feeCap,
			Gas:       gasLimit,
			To:        &to.Address,
			Value:     value,
			Data:      nil,
		})

	signedTx, err := k.SignTx(from, tx, chainid)
	if err != nil {
		return nil, fmt.Errorf("sign transaction: %w", err)
	}

	b, err := tx.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("marshal transaction: %w", err)
	}

	log.Printf("Sending transaction: %s\n", string(b))

	if err := cl.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, fmt.Errorf("send transaction: %w", err)
	}

	txHash := tx.Hash()

	return &txHash, nil
}

func waitTransaction(conn *ethclient.Client, hash common.Hash) error {
	for {
		_, pending, err := conn.TransactionByHash(context.Background(), hash)
		if !pending {
			return nil
		}

		if err != nil {
			return err
		}

		log.Println("Waiting for transaction: " + hash.String())
		time.Sleep(100 * time.Millisecond)
	}
}
