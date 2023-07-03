
- 根据rpc链接和address，返回地址是EOA还是合约

# generate 基本功能
  - go run main.go g w
  - go run main.go g t
  - go run main.go g e

# abi 基础功能
  - go run main.go a -f temp/Lib_AddressManager_abi.json

    打印abi详情

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go a -f temp/Lib_AddressManager_abi.json
    Function name: getAddress
    Input name: _name, type: string
    Output name: , type: address
    State Mutability: view

    Function name: owner
    Output name: , type: address
    State Mutability: view

    Function name: renounceOwnership
    State Mutability: nonpayable

    Function name: setAddress
    Input name: _name, type: string
    Input name: _address, type: address
    State Mutability: nonpayable

    Function name: transferOwnership
    Input name: newOwner, type: address
    State Mutability: nonpayable

    ======================================
    Event name: AddressSet
    Input name: _name, type: string
    Input name: _newAddress, type: address
    Input name: _oldAddress, type: address
    RawName: AddressSet
    Sig: AddressSet(string,address,address)

    Event name: OwnershipTransferred
    Input name: previousOwner, type: address
    Input name: newOwner, type: address
    RawName: OwnershipTransferred
    Sig: OwnershipTransferred(address,address)

    ======================================
    ```

# client 基础功能
  - go run main.go client 

    打印chainID

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go client
    ChainID:  17
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c
    ChainID:  17
    ```

## client account
  - go run main.go c account isContract

    判断是否为合约地址

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a ic 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 是否为合约: false
    ```

  - go run main.go c account balance

    查询balance

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在最新区块的balance: 99984999999972457600528
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b -b 0 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在区块: 0 的balance: 100000000000000000000000
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c a b -b 42 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
    0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 在区块: 42 的balance: 99984999999972457600528
    ```

## client block

333  - go run main.go c block blockNum

    查询最新区块高度

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c bn
    最新区块高度为:  42
    ```

  - go run main.go c block

    查询区块详情

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b 2                                                                    
    根据num: 2 查询到了区块, block.timestamp: 1679394336
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10   
    根据blockhash: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10 查询到了区块, block.timestamp: 1679394336
    ```

  - go run main.go c block header
  
    查询区块头信息

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b h 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    区块 1 的header.ParentHash为: 0xddece7f6aab5880f215b245941a5dfcd2a21d82c52d5aa356f9c0c22ee99babe
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b h 2                                                                 
    区块 2 的header.ParentHash为: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    ```

  - go run main.go c block transactionCount

    查询指定区块的交易数

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c b tc 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10
    根据blockHash: 0x23c1ce310a1d31eb808097c7d6f88f2e5a6e3190df7f28b069f0f352e4116f10 查询到该区块下包含 1 笔txs。
    ```

## client transaction

  - go run main.go client transaction

    查询交易信息

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c tx 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 
    根据txHash: 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 查询到txGas: 2000000
    ```

  - go run main.go client transaction receipt

    查询收据

    ```shell
    (anaconda3)➜  g_bc git:(main) ✗ go run main.go c tx r 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0
    根据txHash: 0x0a781b239c51a92bb3628e61f8e0f93c083ea188ee7c1811659f8d4d20800ff0 查询到Receipt Status: 1
    ```

## client transfer

- go run main.go client transfer SendETHTransaction
  
  ETH transfer

  ```shell
  (anaconda3)➜  g_bc git:(main) ✗ go run main.go c t e -u https://rpc.testnet.mantle.xyz -p 7eefd641410560e690736ee331bd32512c9b58419a877eff2189facbef33cd1e
  txHash:  0xf427e40c59c93b75aa8d634086378fa7c695263eed137f00bc03a8ba326959c3
  ```

- go run main.go client transfer SendERC20Transaction

  ERC20 transfer

```shell

```

- go run main.go rpc smartContract
- go run main.go rpc eventLog
- go run main.go rpc signature
- go run main.go rpc swarm
- go run main.go rpc whisper
- go run main.go rpc utility



# solc

  ## 期望
  generate the ABI from a solidity source file.

  ```shell
  cd packages/contracts/contracts
  rm -rf build
  solc --abi CustomERC20.sol -o build --base-path './' --include-path '../node_modules/'
  solc --bin CustomERC20.sol -o build --base-path './' --include-path '../node_modules/'
  ```

  let's convert the ABI to a Go file that we can import. This new file will contain all the available methods the we can use to interact with the smart contract from our Go application.

  ```shell
  (anaconda3)➜  contracts git:(main) ✗ pwd
  /Users/dl00038ml/PKG/projects/blockChain/a_mantlenetworkio/g_bc/contracts
  (anaconda3)➜  contracts git:(main) ✗ abigen --abi=../packages/contracts/contracts/build/CustomERC20.abi --pkg=CustomERC20 --out=CustomERC20.go
  ```

  ## 现实

  有import的合约, 并不能使用上述的方法生成包含deploy的合约go代码
  `abigen --bin` 无效

  所以使用hardhat deploy contracts
  
  直接使用下述shell
  `bash compile.sh`



# stress

```shell
## 增加l1的ETH
cast rpc "hardhat_setBalance" --rpc-url http://127.0.0.1:9545 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 0x12345678901234567890123456789

# 或者用
curl -X POST -H 'Content-Type: application/json' --data '{"jsonrpc":"2.0","method":"anvil_setBalance","params":["0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", "123456789123456789123456789"],"id":1}' http://localhost:9545 -s |jq

cast b --rpc-url http://127.0.0.1:9545 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266

## 增加l1的MNT
export Proxy__L1MantleToken=0x19C22f181280dF6Ad1d97285cdD430173Df91C12
export deployer=041deb3563e965bce6e803b88b9d25005cb1414c4cdade04181363e87ca9e259
cast send --rpc-url http://127.0.0.1:9545 --private-key $deployer $Proxy__L1MantleToken  "mint(address,uint256)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 12345678901234567890123456789

cast 2d `cast call --rpc-url http://127.0.0.1:9545 0x92aBAD50368175785e4270ca9eFd169c949C4ce1 "balanceOf(address)" 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`

(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s e
ReceiptStatus checking……
L1ERC20 address:  0x8ac5eE52F70AE01dB914bE459D8B3d50126fd6aE
L2ERC20 address:  0x610178dA211FEF7D417bC0e6FeD39F05609AD788
ReceiptStatus checking……
DepostERC20 1237940039285380274899124224
ReceiptStatus checking……


(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s dnt
ReceiptStatus checking……
第 1 次transfer,txHash: 0x8067a508823f4b2c3e6f6408e84027d8137e83946c3f63b7fa29beca74eb6ec2
ReceiptStatus checking……
第 2 次transfer,txHash: 0x280429a0ca5a2ce838e0234a680b72aa47b95fd27836d07ef5ad2ad667408500
ReceiptStatus checking……
第 3 次transfer,txHash: 0xbdc1a2a3b29606e327e4445601c2d87e149254474ea8bef72761ec9e1debe1cc
ReceiptStatus checking……
第 4 次transfer,txHash: 0x89b226549d6884ebb744186c2cd479583be4c811a5cff191dd56b9dc2d8e8313
ReceiptStatus checking……
第 5 次transfer,txHash: 0x01bb0f476461fdba513b40a599d8ee2f7462784ddb221b4a2e939adf1cc62c95
ReceiptStatus checking……
第 6 次transfer,txHash: 0x59a0ea9e136ec186b784733d04dd808689fe05327ab67eeccba3a6349cfd8737
ReceiptStatus checking……
第 7 次transfer,txHash: 0xc3789b024938326ea3f61b463523c81643b3ffe1299574519bf3c0c2eaa337e9
ReceiptStatus checking……
第 8 次transfer,txHash: 0x4b00cc0b9c3c191feea6032ef895040897879369e39de4f887a45a494cbcb2bc
ReceiptStatus checking……
第 9 次transfer,txHash: 0x1879fc1762d0c4c3a4e50ff3c6d2ebaedbab1ad74362418c6060e30d9086191b
ReceiptStatus checking……
第 10 次transfer,txHash: 0xc1aae4b14e297158133b4640cb4b46d9c507bccd95190372843625f3bbe036d2

(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s dnt -l l2
第 1 次transfer,txHash: 0x324c9ac1eec53c2d26ecf4c3e0b04f69a3eb75cd3cf9b32014cd17fa4f72a987
第 2 次transfer,txHash: 0x2acc63e542813fbb1cc7d0a290ded21ac0e08c9e1768b8e9d1bf6f0c0ad3b7ea
第 3 次transfer,txHash: 0xc12e5c52a68aca946a57304bef7f0a0f61febb52df56dbdbf7f70659e9b79e38
第 4 次transfer,txHash: 0x33bf5a204d8c7aa6f2bbf33bde6836a1b4c3d7bc632675254e490a7937796375
第 5 次transfer,txHash: 0x2823de5e2e4c780c3d01bc60c4d12bd065cc0bcaa6bc639d001fd185b98f7570
第 6 次transfer,txHash: 0x8406caca7744dde2b8516adbdfd44f244152c65e45b4b1d0287d52f0fe419369
第 7 次transfer,txHash: 0xcb001150f3e97fda476e2baec16e7fbfe9403c49519f1ee187ea0c89c84e9dc7
第 8 次transfer,txHash: 0x8e46ac46e65129144cb1bc54ae8e75ca2404505652f4f4f0c6766e845eaa6bc5
第 9 次transfer,txHash: 0x7ffac5d5443ee0a169b484c4bad7c3e1e21d7ce613f90f4283cfc9a25764993e
第 10 次transfer,txHash: 0x0092ae4d5254733d912132671dab18094c795544eaf26920334c92354eddc437

(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s d20                      
第 1 次transfer,txHash: 0x777e06c614c801fd2ddbc5320bdf66e1735f262ef00c6379ccfff1d7582e22cb
第 2 次transfer,txHash: 0x47063eec57abae70224b7d151900fed66c744941eb356c0b6891d815b8b66ba0
第 3 次transfer,txHash: 0x54adf63dbc3566fd8022070908f2cf9187b83b83ba8c52cd99f4fa9c1b045d8a
第 4 次transfer,txHash: 0x99515e7f67796b0b5e8ee391b77745d80f57b383323b000baf094fe01268ea13
第 5 次transfer,txHash: 0x78395f0d2101d4c77c4afa30581a01f2f4deaa5958b321fc9a8e4e54ae662c8f
第 6 次transfer,txHash: 0x31fa1e991a1a92ee020698cf80faf94856721b15594f6f002a373fccbae91aa6
第 7 次transfer,txHash: 0x4cdb82fed7b17c63c522522e68324807aef00368baa675e607e76a100aa8e368
第 8 次transfer,txHash: 0x0f68e8712726c05c28fbee9736ea8997cbea1e6a8e339a15617a080653e88afc
第 9 次transfer,txHash: 0xf1e1aae915b435e188ff79272ca916518ac02286c98cc20624beb914c48dcab7
第 10 次transfer,txHash: 0x968233c22f231f2e47ac7578f146a74c3216694310e92be5acc515c48b9d94bc


(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s d20 --isMNT ture
第 1 次transfer,txHash: 0x4c0365387ec6addbc4832114cc3f44090b6b035df68726759668f05a9c6a8a62
第 2 次transfer,txHash: 0x30f195d3a48a19ab9f3f9dd05050292a96bf5484b7e47d2c4277794633872ebd
第 3 次transfer,txHash: 0x07dd2a58b4bbe061b68d2020f24f3f40fb56f032b22c1f46d1b9c2eb0e9ea145
第 4 次transfer,txHash: 0x9275c725a2798077df3affcaf44c23e791ebdd64b720ea31e71e47c6a476d280
第 5 次transfer,txHash: 0xfbe2e8b2649552a3cb5ee4d614830752ed928e61657f797934dc9bcce4827b0f
第 6 次transfer,txHash: 0x073999663e130c2e4a2008a37ea218f7c5278b5dcb73c72fcd594f4baeb11ea5
第 7 次transfer,txHash: 0x7e3b9d5c13ff5a4d3945c333019e6d01aa8b6a23cfb0d3e8c30f332e4d4f5e2a
第 8 次transfer,txHash: 0x85d92000409b5cc30fd792bb4b67b2cc3d58023b20df1f83488cf19f381d67d0
第 9 次transfer,txHash: 0xc5ebe3c664bcc0689badc748c0f3abf9c4fff0e429c259f6efc271a4291b2b1c
第 10 次transfer,txHash: 0xd3d5c59ffa40161ea463e5a0364648784926a7a459740468878c330c5d5b7478

(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s d20 -l l2
第 1 次transfer,txHash: 0x6cd5e0db35ae3993d95e7f535e065f66ea892cfe438df0b5c0fffc0f800c1949
第 2 次transfer,txHash: 0xcf38ea579d75fc17b197419087f47725b3df1be061c307a9a4ae909e383c23d0
第 3 次transfer,txHash: 0xcb829df0f241fdccd7c060b1b9bc260f20158277d9382cc2db646a2756a6addf
第 4 次transfer,txHash: 0x6cc1947589dbaf88f5213dbd77900e6d2da10f447d8d11c0da5d39b3616b970d
第 5 次transfer,txHash: 0x419a0b7e3f7f69fb50c2d60c30c2b49396ea2c0bd10732f7f28b6a9b685b26b7
第 6 次transfer,txHash: 0xc432e0080bdd138c06cf35dba7839d0f4f8f40b93b424d94c30f84afcf4a4482
第 7 次transfer,txHash: 0x1b021fd888bfb4637184a1942655fb2b9a90cfb40553964a45d9b77d377e3eff
第 8 次transfer,txHash: 0xaf287fcd540098260d647ad34de36503ee3a1e6e16d938ade29d54233a4ae8a2
第 9 次transfer,txHash: 0xd8901013701f30101acb57f312d9be600601321cfa68d789256e13826cb3e3c4
第 10 次transfer,txHash: 0x4b0f58856aafb34ccfd5df2572380c7375237cafded567586a779a9c0ceb0611

(anaconda3)➜  g_bc git:(main) ✗ go run main.go m s d20 -l l2 --isETH true
第 1 次transfer,txHash: 0xeeb9cea9c2cd2bd046d764ddc15bf430c883037a447afe655e2f0a0fc94d96f5
第 2 次transfer,txHash: 0x5c64de9088602404d2544c62d652170a6fdeffc31d5f4f2e4cbb98b259334e3f
第 3 次transfer,txHash: 0xc67ac17eac23375f97739ac6087507c969a151a7daf0ab21c7f79527647c3d8d
第 4 次transfer,txHash: 0xf97e8dd423e69617bc8653f88afcf813ed0c10b79280bbfcc5cb1802cf614f20
第 5 次transfer,txHash: 0x10773b679df7cc378b10fb24a6de716b431716e779ca68ae2739942ff12b9ade
第 6 次transfer,txHash: 0x0091045882b2364e7d155244a3b42b986b7b9b1dd76ce5036257531816818e83
第 7 次transfer,txHash: 0xf310ab81998bece9c3c7bc81f8cdf7b2692ef6227ed25c88221e796f449a882b
第 8 次transfer,txHash: 0x217c12cbd9ad252de59c768a455e0ebe4693135a87c45d3521f67d65f7aed4c0
第 9 次transfer,txHash: 0xabae3547b03a937b8e87a88d1262c695a1f26620e3858b2d424fc4ae753dee35
第 10 次transfer,txHash: 0x2729526debdfe1555e51571ea5828e5bdba2a24a086503f61cb8e2a97f0a240e

go run main.go m s st abcdefghij
```