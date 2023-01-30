package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

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

		id, err := client.ChainID(ctx)
		if err != nil {
			log.Println(err)
			client.Close()
			time.Sleep(4 * time.Second)

			continue
		}

		log.Printf("Connected to NetworkID: %d", id.Int64())

		return client
	}
}
