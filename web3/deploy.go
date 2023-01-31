package web3

import (
	"fmt"
	"math/big"

	"golang.org/x/exp/slog"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Davincible/nft-marketplace/contracts"
)

// DeployNFTFactory deploys the factory smart contract.
//
// Optionally set the wait param to true, to require confirmation of contract
// deployment before returning.
func (w3c *W3Connector) DeployNFTFactory(wait bool) (*contracts.KuiperNFTFactory, error) {
	w3c.logger.Debug("Deploying NFTFactory contract")

	addr, tx, nftFactory, err := contracts.DeployKuiperNFTFactory(w3c.auth, w3c.conn)
	if err != nil {
		return nil, fmt.Errorf("deploy NFT Factory: %w", err)
	}

	w3c.logger.Info("Deployed NFTFactory contract", slog.String("addr", addr.Hex()))

	if wait {
		r, err := w3c.WaitForReceipt(tx.Hash())
		if err != nil {
			return nil, err
		}

		w3c.logger.Debug("Received receipt",
			slog.String("contractAddr", r.ContractAddress.Hex()),
			slog.Int64("gasUsed", int64(r.GasUsed)),
		)
	}

	return nftFactory, nil
}

// DeployNFTMarketplace deploys the contract.
//
// Optionally set the wait param to true, to require confirmation of contract
// deployment before returning.
func (w3c *W3Connector) DeployNFTMarketplace(nftFactoryAddr common.Address, wait bool) (*contracts.KuiperNFTMarketplace, error) {
	w3c.logger.Debug("Deploying NFTFactory contract")

	fee := big.NewInt(10) // In percentage

	addr, tx, nftMarketplace, err := contracts.DeployKuiperNFTMarketplace(w3c.auth, w3c.conn, fee, w3c.userAddr, nftFactoryAddr)
	if err != nil {
		return nil, fmt.Errorf("deploy NFT Factory: %w", err)
	}

	w3c.logger.Info("Deployed NFTFactory contract", slog.String("addr", addr.Hex()))

	if wait {
		r, err := w3c.WaitForReceipt(tx.Hash())
		if err != nil {
			return nil, err
		}

		w3c.logger.Debug("Received receipt",
			slog.String("contractAddr", r.ContractAddress.Hex()),
			slog.Int64("gasUsed", int64(r.GasUsed)),
		)
	}

	return nftMarketplace, nil
}
