package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
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

	// Deploy a smart contract and interact with it.
	if err := getGreeting(conn, coinbase); err != nil {
		log.Fatal(err)
	}

	// Transfer BNB across accounts.
	if err := transferToAll(conn, coinbase, transferAmount); err != nil {
		log.Fatal(err)
	}
}
