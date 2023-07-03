/** @type import('hardhat/config').HardhatUserConfig */
require('hardhat-abi-exporter');
require("@nomicfoundation/hardhat-toolbox");
require("@solidstate/hardhat-bytecode-exporter");

module.exports = {
  solidity: "0.8.18",
  abiExporter: {
    path: './abi',
    runOnCompile: true,
    clear: true,
  },
  bytecodeExporter: {
    path: './data',
    runOnCompile: true,
    clear: true,
    flat: true,
    // only: [':*CustomERC20$'],
  },
  networks: {
    hardhat: {},
    local: {
      url: "http://127.0.0.1:9545",
      accounts: ["ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"],
    },
    goerli: {
      url: "https://goerli.infura.io/v3/3f31ba1b9fc54dfd92e1e26a64fba7e0",
      accounts: ["e7bbfa2e0a24248701d7de1745495d8c7e08a835782006db5dddc37442f2ccb2"],
    }
  },
  etherscan: {
    apiKey: "ZMP8FANP6X2Z5H4UA2WVIV6PM82GN77TFE", // Your Etherscan API key
  },
};