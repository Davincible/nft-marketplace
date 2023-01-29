package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	addr = "http://bsc:8545"
)

func main() {
	time.Sleep(30 * time.Second)

	client, err := ethclient.Dial(addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	// addr := "0x0142413a8b0306BaC8127EE16Ec622c1F973ed8c"
	addr := "0x7fd60c817837dcfefca6d0a52a44980d12f70c59"
	account := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(balance) // 25893180161173005034
}
