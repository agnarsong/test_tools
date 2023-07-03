import "@nomiclabs/hardhat-waffle"
import "@openzeppelin/hardhat-upgrades"
import "@openzeppelin/hardhat-defender"
import "@nomiclabs/hardhat-ethers"
import "@typechain/hardhat"
import "hardhat-gas-reporter"
import "hardhat-contract-sizer"
import "hardhat-abi-exporter"
import "./tasks/swap"
import "./tasks/upgrade"
import "./tasks/swapSubtask"
require('dotenv').config()

module.exports = {
    defaultNetwork: 'hardhat',
    defender: {
        apiKey: "[apiKey]",
        apiSecret: "[apiSecret]",
    },
    networks: {
        hardhat: {
            allowUnlimitedContractSize: true,
        },
        qaNew: {
            url: 'http://10.41.20.51:8545',
            gasPrice: 5000000000,
            chainId: 7001,
            gas: 4100000,
            accounts:['e61a127ec0e87c9eb19e02ff2733ec48c614076bd6cde26ca73ef0f00fe546b4']
        },
        testl2: { // Optimism Goerli
            //url: "https://mantle-verifier-testnet.qa.davionlabs.com",
            url: "https://rpc.testnet.mantle.xyz/",
            //url: "http://10.45.25.101:8545",
            accounts: ['7eefd641410560e690736ee331bd32512c9b58419a877eff2189facbef33cd1e'],
            chainId: 5001,
            gas: 10000000,
            gasPrice: 1,
        },
        arbitrum: {
            url: 'https://rinkeby.arbitrum.io/rpc',
            gasPrice: 200000000,
            chainId: 421611,
            gas: 5000000,
            // accounts:['96e8e32341ce890aff8b46066f7b77a6d2ab2115a24c365e9de1fbed49e04837']
            accounts: [process.env.PRI_KEY]
        },
        localhost: {
            url: "http://127.0.0.1:8545",
            accounts: [process.env.PRI_KEY]
        },
        g:{ // Goerli
            url: "https://goerli.davionlabs.com",
            accounts: [process.env.PRI_KEY],
            chainId: 5
            //explorer: https://goerli.etherscan.io
        },
        qal2: { // Optimism Goerli
            //url: "https://mantle-verifier.qa.davionlabs.com",
            url: "https://mantle-l2geth.qa.davionlabs.com",
            accounts: ['0x267a3f27ebf78887da180d8bd9ed7cff4b80e7227220db64eac0a3366117f83e'],
            chainId: 1705003,
            gas: 10000000,
            gasPrice: 1,
        },
        teleport: {
            url: 'https://teleport-localvalidator.qa.davionlabs.com/',
            gasPrice: 5000000000,
            chainId: 7001,
            gas: 4100000,
            // accounts:['24ad33fb88a6c2347ec90178c881969e59571c6ad8cc0f597e7c7d87354df3f8']
            //usdt
            //accounts:['96e8e32341ce890aff8b46066f7b77a6d2ab2115a24c365e9de1fbed49e04837']
            //qa 发钱test net admin私钥
            //accounts:['2307796387358a2a99fbbb88312dc6516ed7ab02bd8f04cc44019e4818560157']
            accounts: [process.env.PRI_KEY]
        },
        rinkeby: {
            url: 'https://rinkeby.infura.io/v3/023f2af0f670457d9c4ea9cb524f0810',
            gasPrice: 1500000000,
            chainId: 4,
            gas: 5000000,
            // accounts: [process.env.PRI_KEY]
            accounts:['0x7eefd641410560e690736ee331bd32512c9b58419a877eff2189facbef33cd1e']
            //accounts:['8f14df1da1a318bec99800b72c5031e4fdc4ec017f00ab9659339ecb0193120e']
        },
        bsctest: {
            url: 'https://data-seed-prebsc-2-s1.binance.org:8545',
            gasPrice: 10000000000,
            chainId: 97,
            gas: 5000000,
            //admin adder
            // accounts:['96e8e32341ce890aff8b46066f7b77a6d2ab2115a24c365e9de1fbed49e04837']
            //accounts:['8f14df1da1a318bec99800b72c5031e4fdc4ec017f00ab9659339ecb0193120e']
            accounts: [process.env.PRI_KEY]
        },
        opg: { // Optimism Goerli
            url: "https://goerli.optimism.io",
            accounts: [process.env.PRI_KEY],
            chainId: 420,
            gas: 10000000,
            gasPrice: 5000000,
            //explorer: https://blockscout.com/optimism/goerli/
        },
        zkl1: {
            url: "http://10.45.21.31:8585",
            accounts: [process.env.PRI_KEY],
            chainId: 9,
            gas: 10000000,
            gasPrice: 5000000,
            //explorer: https://blockscout.com/optimism/goerli/
        },
        zkl2: {
            url: "http://10.45.21.31:8050",
            accounts: [process.env.PRI_KEY],
            chainId: 270,
            // gas: 10000000,
            // gasPrice: 5000000,
            //explorer: https://blockscout.com/optimism/goerli/
        },
        pll1: {
            url: "http://10.45.21.31:8545",
            accounts: [process.env.PRI_KEY],
            chainId: 1337,
            gas: 10000000,
            gasPrice: 5000000,
            //explorer: https://blockscout.com/optimism/goerli/
        },
        pll2: {
            url: "http://10.45.21.31:8123",
            accounts: [process.env.PRI_KEY],
            chainId: 1001,
            gas: 2100000,
            gasPrice: 50000000
            //explorer: https://blockscout.com/optimism/goerli/
        },
        locall2: {
            url: "http://127.0.0.1:8545",
            // gas: 210000,
            // gasPrice: 1,
            accounts: [//dev 环境专用
                'ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
                //'7eefd641410560e690736ee331bd32512c9b58419a877eff2189facbef33cd1e',
            ],
        },
        locall1: {
            chainId: 31337,
            url: "http://127.0.0.1:9545",
            gas: 4100000,
            gasPrice: 5000000000,
            accounts: [//dev 环境专用
                '041deb3563e965bce6e803b88b9d25005cb1414c4cdade04181363e87ca9e259',
            ],
        },
    },
    solidity: {
        compilers: [
            {
                version: '0.6.6',
                settings: {
                    optimizer: {
                        enabled: true,
                        runs: 5,
                    },
                }
            },
            {
                version: '0.6.12',
            },
            {
                version: '0.4.17',
            }
        ],
        settings: {
            optimizer: {
                enabled: true,
                runs: 200,
            },
        }
    },
    gasReporter: {
        enabled: true,
        showMethodSig: true,
        maxMethodDiff: 10,
    },
    contractSizer: {
        alphaSort: true,
        runOnCompile: true,
        disambiguatePaths: false,
    },
    abiExporter: {
        path: './build',
        runOnCompile: true,
        clear: true,
        flat: true,
        except: ['@openzeppelin/contracts'],
        spacing: 2
        // pretty: true,
    }
}