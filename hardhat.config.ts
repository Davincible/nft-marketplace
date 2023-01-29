import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-deploy";
// import "@nomiclabs/hardhat-ethers";

let mnemonic = process.env.MNEMONIC;
if (!mnemonic) {
  // FOR DEV ONLY, SET IT IN .env files if you want to keep it private
  // (IT IS IMPORTANT TO HAVE A NON RANDOM MNEMONIC SO THAT SCRIPTS CAN ACT ON THE SAME ACCOUNTS)
  mnemonic = "test test test test test test test test test test test junk";
}

const config: HardhatUserConfig = {
  // solidity: "0.8.17",
  solidity: {
    compilers: [
      {
        version: "0.8.17",
        settings: {
          optimizer: {
            enabled: true,
            runs: 2000,
          },
        },
      },
    ],
  },
  namedAccounts: {
    deployer: {
      default: 0, // here this will by default take the first account as deployer
    },
  },
  defaultNetwork: "localhost",
  networks: {
    localhost: {
      url: "http://localhost:8545",
      /*
        notice no mnemonic here? it will just use account 0 of the hardhat node to deploy
        (you can put in a mnemonic here to set the deployer locally)
      */
    },
    rinkeby: {
      url: `https://rinkeby.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic },
    },
    mainnet: {
      url: `https://mainnet.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic },
    },
    ropsten: {
      url: `https://ropsten.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic },
    },
    goerli: {
      url: `https://goerli.infura.io/v3/${process.env.INFURA_TOKEN}`,
      accounts: { mnemonic },
    },
    xdai: {
      url: "https://dai.poa.network",
      chainId: 100,
      gas: 500000,
      gasPrice: 1000000000,
      accounts: { mnemonic },
    },
  },
};

export default config;
