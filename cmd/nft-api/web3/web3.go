package web3

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slog"
)

var (
	defaultInitTimeout    = 180 * time.Second
	defaultTimeout        = 15 * time.Second
	defaultReceiptTimeout = 60 * time.Second
)

// Errors.
var (
	ErrConnectionFailed = errors.New("failed to establish initial connection")
	ErrNoKeystoreDir    = errors.New("no keystore directory path provided")
	ErrEmptyKeystore    = errors.New("keystore contains no accounts, please add accounts")
	ErrNoAccountAddr    = errors.New("no accounts address set")
	ErrNoAddress        = errors.New("no blockchain network address provided")
	ErrReceiptDeadline  = errors.New("no receipt available within deadline")
)

// W3Connector object provides methods to interact with any JSON-RPC capable
// blockchain.
type W3Connector struct {
	conn     *ethclient.Client
	keystore *keystore.KeyStore
	logger   *slog.Logger

	// userAddr is the publickey address of the account to use to deploy and
	// interact with contracts.
	userAddr common.Address

	// chainID is the ID of the network the client is connected to. This is used
	// when sigining transactions to make sure that tranascations cannot be replayed
	// on other blockchain networks with a different chain ID.
	chainID *big.Int

	// auth is the object that signs transactions with private key from the address
	// as set in userAddr.
	auth *bind.TransactOpts

	// initTimeout is the timeout used when initially connecting to the network.
	// When no successful connection has been established within this time duration
	// the connector will return an error.
	initTimeout time.Duration

	// timeout used for calls to the blockchain network.
	timeout time.Duration

	// receiptTimeout is the max duration to wait for a transaction block to be
	// mined, and for the transaction receipt to be available.
	receiptTimeout time.Duration
}

// NewW3Connector connects to a blockchain and returns a connector with a client.
func NewW3Connector(logger *slog.Logger, addr, keystore, userAddr string) (*W3Connector, error) {
	if len(addr) == 0 {
		return nil, ErrNoAddress
	}

	if len(userAddr) == 0 {
		return nil, ErrNoAccountAddr
	}

	conn, err := ethclient.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial address '%s': %w", addr, err)
	}

	w3c := W3Connector{
		conn:        conn,
		timeout:     defaultTimeout,
		userAddr:    common.HexToAddress(userAddr),
		initTimeout: defaultInitTimeout,
		logger: logger.With(
			slog.String("addr", addr),
			slog.String("userAddr", userAddr),
		),
	}

	if err := w3c.connect(); err != nil {
		return nil, err
	}

	if err := w3c.openKeystore(keystore); err != nil {
		return nil, err
	}

	return &w3c, nil
}

// Close the client connection.
func (w3c *W3Connector) Close() {
	w3c.conn.Close()
}

// connect will make the first request to the web3 endpoint, to make sure it is
// online. If no connection has been successful within the timelimit it will
// return an error.
//
// It will also set the ChainID.
func (w3c *W3Connector) connect() error {
	deadline := time.After(defaultInitTimeout)

	w3c.logger.Debug("Connecting to network")

	for {
		select {
		case <-deadline:
			return ErrConnectionFailed
		default:
			ctx, cancel := context.WithTimeout(context.Background(), w3c.timeout)
			defer cancel()

			chainID, err := w3c.conn.ChainID(ctx)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			w3c.chainID = chainID

			w3c.logger.Info("Connected to network", slog.Int64("chainID", w3c.chainID.Int64()))

			return nil
		}
	}
}

// openKeystore opens the keystore directory containing addresses with private
// keys in the web3 account format.
func (w3c *W3Connector) openKeystore(path string) error {
	if len(path) == 0 {
		return ErrNoKeystoreDir
	}

	w3c.keystore = keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)

	if len(w3c.keystore.Accounts()) == 0 {
		return ErrEmptyKeystore
	}

	acc := accounts.Account{Address: w3c.userAddr}

	if err := w3c.keystore.Unlock(acc, ""); err != nil {
		return fmt.Errorf("unlock account: %w", err)
	}

	auth, err := bind.NewKeyStoreTransactorWithChainID(w3c.keystore, acc, w3c.chainID)
	if err != nil {
		return fmt.Errorf("create auth transactor: %w", err)
	}

	w3c.auth = auth

	return nil
}

func (w3c *W3Connector) WaitForReceipt(addr common.Hash) (*types.Receipt, error) {
	deadline := time.After(defaultInitTimeout)

	w3c.logger.Debug("Waiting for transaction receipt", slog.String("addr", addr.Hex()))

	for {
		select {
		case <-deadline:
			return nil, ErrReceiptDeadline
		default:
			ctx, cancel := context.WithTimeout(context.Background(), w3c.timeout)
			defer cancel()

			r, err := w3c.conn.TransactionReceipt(ctx, addr)
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
