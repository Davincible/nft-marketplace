package web3

import (
	"log"
	"os"
	"testing"

	"golang.org/x/exp/slog"

	"github.com/stretchr/testify/assert"

	"github.com/Davincible/nft-marketplace/utils/env"
)

var (
	connector *W3Connector

	userAddr    = env.GetEnv("COINBASE_ADDR", "0x7fd60C817837dCFEFCa6D0A52A44980d12F70C59")
	networkAddr = env.GetEnv("NETWORK_ADDR", "http://localhost:8545")
	keystoreDir = env.GetEnv("KEYSTORE_DIR", "./accounts")
)

func init() {
	logger := slog.New(
		slog.HandlerOptions{Level: slog.LevelDebug}.NewTextHandler(os.Stdout),
	)

	c, err := NewW3Connector(logger, networkAddr, keystoreDir, userAddr)
	if err != nil {
		log.Fatal(err)
	}

	connector = c
}

func TestDeployNFTFactory(t *testing.T) {
	_, err := connector.DeployNFTFactory(true)
	assert.NoError(t, err)
}
