# Sample NFT Market API

> WIP: This project is a work in progress.

This project will be a sample NFT marketplace with HTTP API, through Go bindings
interacting with the JSON-RPC API of any EVM compatible blockchain network.

This project uses a Binance Smart Chain (BSC) local network for development, with
the [KuiperNFT Marketplace](https://github.com/kofkuiper/kuiper-nft-marketplace).

## Running the BSC Network

A local BSC network of 2 nodes running the BSC version of `geth` with the parlia
consencus protocol will be spun up.

All accounts available in the network can be found in the `accounts` folder.

```bash
# Step 1: Build Docker containers. This takes a couple of minutes.
# Step 2: Bootstrap tne BSC network. Also run this if you want to reset the network.
# Step 3: Run the docker containers.
$ make docker-build bsc-bootstrap docker-compose
```

## Running tests

```bash
# This runs a small progrom to demonstrate the usage of Go with web3.
$ ADDRESS="http://localhost:8545" go run cmd/bsc-init/*.go

# This runs the tests so far available.
$ KEYSTORE_DIR="$(pwd)/accounts" go test -v ./cmd/nft-api/...
```

## Dev commands

```bash
# Install Solidity dependencies.
$ make init

# Compile all Solidity smart contracts.
$ make compile-contracts

# Generate Go bindings for smart contracts.
$ make gen-bindings
```


