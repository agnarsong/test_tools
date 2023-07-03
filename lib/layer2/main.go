package layer2

import (
	"bufio"
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"mantle/test/lib"
	"mantle/test/lib/layer2/bindings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
)

type AddressList struct {
	AddressDictator                     string `json:"AddressDictator"`
	BVML1CrossDomainMessenger           string `json:"BVM_L1CrossDomainMessenger"`
	BondManager                         string `json:"BondManager"`
	CanonicalTransactionChain           string `json:"CanonicalTransactionChain"`
	ChainStorageContainerCTCBatches     string `json:"ChainStorageContainer-CTC-batches"`
	ChainStorageContainerSCCBatches     string `json:"ChainStorageContainer-SCC-batches"`
	ChugSplashDictator                  string `json:"ChugSplashDictator"`
	L1StandardBridgeForVerificationOnly string `json:"L1StandardBridge_for_verification_only"`
	LibAddressManager                   string `json:"Lib_AddressManager"`
	ProxyBVML1CrossDomainMessenger      string `json:"Proxy__BVM_L1CrossDomainMessenger"`
	ProxyBVML1StandardBridge            string `json:"Proxy__BVM_L1StandardBridge"`
	ProxyTSSGroupManager                string `json:"Proxy__TSS_GroupManager"`
	ProxyTSSStakingSlashing             string `json:"Proxy__TSS_StakingSlashing"`
	StateCommitmentChain                string `json:"StateCommitmentChain"`
	TssGroupManager                     string `json:"TssGroupManager"`
	TssStakingSlashing                  string `json:"TssStakingSlashing"`
	AddressManager                      string `json:"AddressManager"`
	Proxy__L1MantleToken                string `json:"Proxy__L1MantleToken"`
	Proxy__Verifier                     string `json:"Proxy__Verifier"`
	BVM_EigenDataLayrFee                string `json:"BVM_EigenDataLayrFee"`
	Rollup                              string `json:"Rollup"`
	Proxy__BVM_EigenDataLayrFee         string `json:"Proxy__BVM_EigenDataLayrFee"`
	VerifierEntry                       string `json:"VerifierEntry"`
	BVM_EigenDataLayrChain              string `json:"BVM_EigenDataLayrChain"`
	Proxy__BVM_EigenDataLayrChain       string `json:"Proxy__BVM_EigenDataLayrChain"`
	L1MantleToken                       string `json:"L1MantleToken"`
	AssertionMap                        string `json:"AssertionMap"`
	Proxy__Rollup                       string `json:"Proxy__Rollup"`
}

type Rpc struct {
	URL     string `json:"url"`
	ChainID int    `json:"chainID"`
}

type StressEnv struct {
	L1URL                   string      `json:"l1url"`
	L1WS                    string      `json:"l1ws"`
	L1ChainID               int         `json:"l1chainid"`
	L2URL                   string      `json:"l2url"`
	L2WS                    string      `json:"l2ws"`
	RpcList                 []Rpc       `json:"rpclist"`
	ToAddress               string      `json:"toaddress"`
	TokenAddress            string      `json:"tokenaddress"`
	L2MNTAddress            string      `json:"l2mntaddress"`
	L2EthAddress            string      `json:"l2ethaddress"`
	L2BridgeAddress         string      `json:"l2bridgeaddress"`
	EnvType                 string      `json:"envtype"`
	PrivateKeyList          [][]string  `json:"privatekeylist"`
	DeployBool              bool        `json:"deploybool"`
	L1ERC20Address          string      `json:"l1erc20address"`
	L2ERC20Address          string      `json:"l2erc20address"`
	L1GasLimit              uint64      `json:"l1gaslimit"`
	L2GasLimit              uint64      `json:"l2gaslimit"`
	Amount                  string      `json:"amount"`
	AddressList             AddressList `json:"addresslist"`
	L2cdmAddress            string      `json:"l2cdmaddress"`
	BVM_L2ToL1MessagePasser string      `json:"bvm_l2tol1messagepasser"`
}

type MantleCenter struct {
	Env             StressEnv
	L1Client        *ethclient.Client
	L1WSClient      *ethclient.Client
	L2Client        *ethclient.Client
	L2WSClient      *ethclient.Client
	UserPrivateKey  *ecdsa.PrivateKey
	UserAddress     common.Address
	ToAddress       common.Address
	L1ERC20         *bindings.L1CustomERC20
	L2ERC20         *bindings.L2CustomERC20
	L1ERC20Address  common.Address
	L2ERC20Address  common.Address
	L1MNTAddress    common.Address
	L2MNTAddress    common.Address
	L2ETHAddress    common.Address
	L1Bridge        *bindings.L1StandardBridge
	L2Bridge        *bindings.L2StandardBridge
	L1BridgeAddress common.Address
	L2BridgeAddress common.Address
	SCCAddress      common.Address
	SCC             *bindings.StateCommitmentChain
	CTCAddress      common.Address
	CTC             *bindings.CanonicalTransactionChain
}

type StateRootBatch struct {
	BlockNumber uint64
	StateRoots  []string
	Header      struct {
		BatchIndex        *big.Int
		BatchRoot         [32]byte
		BatchSize         *big.Int
		PrevTotalElements *big.Int
		Signature         []byte
		extraData         []byte
	}
}

func InitSc(mc *MantleCenter) (err error) {

	mc.L1Client, err = lib.NewEthClient(mc.Env.L1URL)
	if err != nil {
		return fmt.Errorf("mc.L1Client err: %v", err)
	}

	if mc.Env.L1WS != "" {
		mc.L1WSClient, err = lib.NewEthClient(mc.Env.L1WS)
		if err != nil {
			return fmt.Errorf("mc.L1WSClient err: %v", err)
		}
	}

	mc.L2Client, err = lib.NewEthClient(mc.Env.L2URL)
	if err != nil {
		return fmt.Errorf("mc.L2Client err: %v", err)
	}

	if mc.Env.L2WS != "" {
		mc.L2WSClient, err = lib.NewEthClient(mc.Env.L2WS)
		if err != nil {
			return fmt.Errorf("mc.L2Client err: %v", err)
		}
	}

	mc.UserPrivateKey, err = crypto.HexToECDSA(mc.Env.PrivateKeyList[0][0])
	if err != nil {
		return fmt.Errorf("mc.UserPrivateKey err: %v", err)
	}
	mc.UserAddress = crypto.PubkeyToAddress(mc.UserPrivateKey.PublicKey)

	var addressURL string
	switch mc.Env.EnvType {
	case "local":
		addressURL = mc.Env.L1URL[:len(mc.Env.L1URL)-4] + "8081/addresses.json"
	case "qa":
		addressURL = ""
	}

	if addressURL != "" {
		// 请求address.json
		resp, err := http.Get(addressURL)
		if err != nil {
			return fmt.Errorf("http.Get err: %v", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("ioutil.ReadAll err: %v", err)
		}

		var addressList AddressList
		if err = json.Unmarshal([]byte(body), &addressList); err != nil {
			return fmt.Errorf("json.Unmarshal err: %v", err)
		}

		mc.Env.AddressList = addressList
	}
	mc.L1ERC20Address = common.HexToAddress(mc.Env.L1ERC20Address)
	mc.L2ERC20Address = common.HexToAddress(mc.Env.L2ERC20Address)
	mc.L1MNTAddress = common.HexToAddress(mc.Env.AddressList.Proxy__L1MantleToken)
	mc.L2MNTAddress = common.HexToAddress(mc.Env.L2MNTAddress)
	mc.L2ETHAddress = common.HexToAddress(mc.Env.L2EthAddress)

	mc.SCCAddress = common.HexToAddress(mc.Env.AddressList.StateCommitmentChain)
	mc.SCC, err = bindings.NewStateCommitmentChain(mc.SCCAddress, mc.L1Client)
	if err != nil {
		return fmt.Errorf("bindings.NewStateCommitmentChain err: %v", err)
	}

	mc.CTCAddress = common.HexToAddress(mc.Env.AddressList.CanonicalTransactionChain)
	mc.CTC, err = bindings.NewCanonicalTransactionChain(mc.CTCAddress, mc.L1Client)
	if err != nil {
		return fmt.Errorf("bindings.NewChainStorageContainer err: %v", err)
	}

	mc.L1BridgeAddress = common.HexToAddress(mc.Env.AddressList.ProxyBVML1StandardBridge)
	mc.L1Bridge, err = bindings.NewL1StandardBridge(mc.L1BridgeAddress, mc.L1Client)
	if err != nil {
		return fmt.Errorf("bindings.NewL1StandardBridge err: %v", err)
	}

	mc.L2BridgeAddress = common.HexToAddress(mc.Env.L2BridgeAddress)
	mc.L2Bridge, err = bindings.NewL2StandardBridge(mc.L2BridgeAddress, mc.L2Client)
	if err != nil {
		return fmt.Errorf("bindings.NewL2StandardBridge err: %v", err)
	}

	// 打开包含私钥和地址的文件
	file, err := os.Open("wallets.csv")
	if err != nil {
		fmt.Println("os.Open err: ", err)
	}
	defer file.Close()
	// 从文件中读取每个钱包的私钥和地址，并为每个钱包创建一笔交易
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("reader.ReadAll err: ", err)
	}
	for _, r := range record {
		mc.Env.PrivateKeyList = append(mc.Env.PrivateKeyList, r)
	}

	return nil
}

// Distribute Native Token
func DNT(mc *MantleCenter, layer string) error {

	c := mc.L1Client
	if layer == "l2" {
		c = mc.L2Client
	}

	//文件读取
	f, err := os.Open("./wallets.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// 多协程初始化数据
	for i := 0; i < len(preData); i++ {

		tx, cid, privateKey, err := lib.SignETHTx1(
			c,
			mc.Env.PrivateKeyList[0][0],
			preData[i][1],
			mc.Env.Amount,
			[]byte(""),
		)
		if err != nil {
			return err
		}

		signedTx, err := lib.SignETHTx2(tx, cid, privateKey)
		if err != nil {
			return err
		}

		txHash, err := lib.SendTransaction(c, signedTx)
		if err != nil {
			return err
		}
		fmt.Printf("第 %v 次transfer,txHash: %v\n", i+1, txHash)

		time.Sleep(time.Microsecond * time.Duration(400))

	}

	return nil
}

func DeployL1CustomERC20(mc *MantleCenter) error {

	auth, err := lib.GetAuth(mc.L1Client, mc.Env.PrivateKeyList[0][0])
	if err != nil {
		return fmt.Errorf("lib.GetAuth err: %v", err)
	}

	var tx *types.Transaction
	mc.L1ERC20Address, tx, mc.L1ERC20, err = bindings.DeployL1CustomERC20(
		auth, mc.L1Client, "L1CustomERC20", "L1E")
	if err != nil {
		return fmt.Errorf("bindings.DeployL1CustomERC20 err: %v", err)
	}

	err = lib.CheckReceiptStatus(mc.L1Client, tx.Hash())
	if err != nil {
		return fmt.Errorf("lib.GetReStatus err: %v", err)
	}

	mc.Env.L1ERC20Address = mc.L1ERC20Address.Hex()

	return nil
}

func DeployL2CustomERC20(mc *MantleCenter) error {

	auth, err := lib.GetAuth(mc.L2Client, mc.Env.PrivateKeyList[0][0])
	if err != nil {
		return fmt.Errorf("lib.GetAuth err: %v", err)
	}

	var tx *types.Transaction
	mc.L2ERC20Address, tx, mc.L2ERC20, err = bindings.DeployL2CustomERC20(
		auth,
		mc.L2Client,
		mc.L2BridgeAddress,
		mc.L1ERC20Address,
	)
	if err != nil {
		return fmt.Errorf("bindings.DeployL2CustomERC20 err: %v", err)
	}

	err = lib.CheckReceiptStatus(mc.L2Client, tx.Hash())
	if err != nil {
		return fmt.Errorf("lib.GetReStatus err: %v", err)
	}

	mc.Env.L2ERC20Address = mc.L2ERC20Address.Hex()

	return nil
}

// Distribute ERC20
func D20(c *ethclient.Client, tokenAddress common.Address, pri string, amount string) error {

	//文件读取
	f, err := os.Open("./wallets.csv")
	if err != nil {
		return err
	}
	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	if err != nil {
		return err
	}

	amountB, b := lib.ParseAmount(amount)
	if !b {
		return fmt.Errorf("ParseAmount return false")
	}

	// 多协程初始化数据
	for i := 0; i < len(preData); i++ {

		tx, err := lib.TransferERC20(c,
			tokenAddress,
			pri,
			common.HexToAddress(preData[i][1]),
			amountB,
		)
		if err != nil {
			return err
		}

		fmt.Printf("第 %v 次transfer,txHash: %v\n", i+1, tx.Hash())

		time.Sleep(time.Microsecond * time.Duration(400))
	}

	return nil
}

// deposit
func Deposit(mc *MantleCenter, prk string, l1token common.Address, l2token common.Address,
	amount *big.Int) (tx *types.Transaction, err error) {

	auth, err := lib.GetAuth(mc.L1Client, prk)
	if err != nil {
		return nil, fmt.Errorf("GetAuth err: %v", err)
	}

	// 如果tokenAddress 不等于 0 地址
	if l1token != common.HexToAddress("0") {

		tx, err = lib.ApproveERC20(mc.L1Client, l1token,
			prk, mc.L1BridgeAddress, amount)
		if err != nil {
			return nil, fmt.Errorf("ApproveERC20 err: %v", err)
		}

		err = lib.CheckReceiptStatus(mc.L1Client, tx.Hash())
		if err != nil {
			return nil, fmt.Errorf("GetReStatus err: %v", err)
		}

		auth.Nonce.Add(auth.Nonce, common.Big1)
		fmt.Println("DepostERC20", amount)
		tx, err = mc.L1Bridge.DepositERC20(auth,
			l1token,
			l2token,
			amount,
			2_000_000,
			[]byte("0x"),
		)
		// tx, err = DepostERC20(mc.L1Client, mc.L1BridgeAddress, prk, l1token, l2token, amount, big.NewInt(int64(mc.Env.L2GasLimit)))
		if err != nil {
			return nil, fmt.Errorf("DepositERC20 err: %v", err)
		}
		if err = lib.CheckReceiptStatus(mc.L1Client, tx.Hash()); err != nil {
			return nil, fmt.Errorf("txHash: %v, \nDepostERC20 GetReStatus err: %v", tx.Hash(), err)
		}

		return tx, err
	}

	fmt.Println("DepostETH", amount)
	auth.Value = amount
	tx, err = mc.L1Bridge.DepositETH(auth,
		15_000_000,
		[]byte("0x"),
	)
	// tx, err = DepostETH(mc.L1Client, mc.L1BridgeAddress, prk, amount)
	if err != nil {
		return nil, fmt.Errorf("DepostETH err: %v", err)
	}

	if err = lib.CheckReceiptStatus(mc.L1Client, tx.Hash()); err != nil {
		return nil, fmt.Errorf("txHash: %v, \nDepostETH GetReStatus err: %v", tx.Hash(), err)
	}

	return tx, err
}

func DepostETH(c *ethclient.Client, tokenAddress common.Address,
	pri string, amount *big.Int) (*types.Transaction, error) {

	privateKey, _, fromAddress, err := lib.AnalysePrivateKey(pri)
	if err != nil {
		return nil, err
	}

	transferFnSignature := []byte("depositETH(uint32,bytes)")
	h := sha3.NewLegacyKeccak256()
	h.Write(transferFnSignature)
	methodID := h.Sum(nil)[:4]

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	paddedData := common.LeftPadBytes(common.Hex2Bytes(""), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAmount...)
	data = append(data, paddedData...)
	fmt.Println("data: ", common.Bytes2Hex(data))
	gasLimit, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	nonce, gasPrice, chainID, err := lib.QBasic(c, fromAddress)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func DepostERC20(c *ethclient.Client, tokenAddress common.Address, pri string,
	l1token common.Address, l2token common.Address, amount *big.Int,
	l2Gas *big.Int) (*types.Transaction, error) {

	privateKey, _, fromAddress, err := lib.AnalysePrivateKey(pri)
	if err != nil {
		return nil, err
	}

	// abi 中的args间不能有空格
	abi := "depositERC20(address, address, uint256, uint32, bytes)"
	abi = strings.ReplaceAll(abi, " ", "")
	methodID := crypto.Keccak256([]byte(abi))[:4]

	paddedL1Token := common.LeftPadBytes(l1token.Bytes(), 32)
	paddedL2Token := common.LeftPadBytes(l2token.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	paddedL2Gas := common.LeftPadBytes(big.NewInt(2000000).Bytes(), 32)
	// 当bytes 类型输入为空的时候, 目前除了的有问题
	// except 0x58a997f6000000000000000000000000457ccf29090fe5a24c19c1bc95f492168c0eafdb000000000000000000000000a513e6e4b8f2a923d98304ec87f64353c4d5c853000000000000000000000000000000000000000000000000000000000000303900000000000000000000000000000000000000000000000000000000001e848000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000000
	// actual 0x58a997f6000000000000000000000000457ccf29090fe5a24c19c1bc95f492168c0eafdb000000000000000000000000a513e6e4b8f2a923d98304ec87f64353c4d5c853000000000000000000000000000000000000000000000000000000000000303900000000000000000000000000000000000000000000000000000000001e848000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	// 存在一个 a 的出入
	paddedData := common.LeftPadBytes([]byte(""), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedL1Token...)
	data = append(data, paddedL2Token...)
	data = append(data, paddedAmount...)
	data = append(data, paddedL2Gas...)
	data = append(data, paddedData...)
	fmt.Println("data: ", common.Bytes2Hex(data))
	gasLimit, err := c.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		return nil, err
	}

	nonce, gasPrice, chainID, err := lib.QBasic(c, fromAddress)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func Withdraw(mc *MantleCenter, prk string, l2token common.Address,
	amount *big.Int) (*types.Transaction, error) {

	auth, err := lib.GetAuth(mc.L2Client, prk)
	if err != nil {
		return nil, fmt.Errorf("GetAuth err: %v", err)
	}

	tx, err := mc.L2Bridge.Withdraw(auth,
		l2token,
		amount,
		3000_000,
		[]byte("0x"),
	)
	if err != nil {
		return nil, err
	}

	if err = lib.CheckReceiptStatus(mc.L2Client, tx.Hash()); err != nil {
		return tx, fmt.Errorf("Withdraw GetReStatus err: %v", err)
	}
	return tx, nil
}

func BETH(mc *MantleCenter, holder common.Address, blockNumber *big.Int) (l1b *big.Int, l2b *big.Int, err error) {

	l1b, err = mc.L1Client.BalanceAt(context.Background(), holder, blockNumber)
	if err != nil {
		return l1b, l2b, fmt.Errorf("L1Client.BalanceAt err: %v", err)
	}

	l2b, err = lib.BalanceOf(mc.L2Client, mc.L2ETHAddress, holder, blockNumber)
	if err != nil {
		return l1b, l2b, fmt.Errorf("L2Client.BalanceOf err: %v", err)
	}

	return l1b, l2b, nil
}

func BERC20(mc *MantleCenter, holder common.Address, l1token common.Address,
	l2token common.Address, blockNumber *big.Int) (l1b *big.Int, l2b *big.Int, err error) {

	l1b, err = lib.BalanceOf(mc.L1Client, l1token, holder, blockNumber)
	if err != nil {
		return l1b, l2b, fmt.Errorf("L2Client.BalanceOf err: %v", err)
	}

	l2b, err = lib.BalanceOf(mc.L2Client, l2token, holder, blockNumber)
	if err != nil {
		return l1b, l2b, fmt.Errorf("L2Client.BalanceOf err: %v", err)
	}

	return l1b, l2b, nil
}

func GetStateBatchAppendedEventByBatchIndex(
	mc *MantleCenter,
	batchIndex []*big.Int,
	start, end uint64) (
	event *bindings.StateCommitmentChainStateBatchAppended,
	err error) {

	sccFilter, err := bindings.NewStateCommitmentChainFilterer(mc.SCCAddress, mc.L1Client)
	if err != nil {
		return nil, fmt.Errorf("NewStateCommitmentChainFilterer err: %v", err)
	}

	filterOpts := bind.FilterOpts{
		Start:   start,
		End:     &end,
		Context: context.Background(),
	}

	iterator, err := sccFilter.FilterStateBatchAppended(&filterOpts, batchIndex)
	if err != nil {
		return nil, fmt.Errorf("sccFilter.FilterStateBatchAppended err: %v", err)
	}

	// 迭代所有事件
	for iterator.Next() {
		// 从迭代器中获取事件对象
		event = iterator.Event
	}

	if err := iterator.Error(); err != nil {
		// 处理迭代器错误
		return nil, fmt.Errorf("iterator err: %v", err)
	}

	return event, nil
}

func GetStateBatchAppendedEventByTransactionIndex(
	mc *MantleCenter,
	transactionIndex uint64,
) (*bindings.StateCommitmentChainStateBatchAppended, error) {

	totalBatches, err := mc.SCC.GetTotalBatches(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("GetTotalBatches err: %v", err)
	}

	isEventHi := func(event *bindings.StateCommitmentChainStateBatchAppended, index uint64) bool {
		prevTotalElements := event.PrevTotalElements.Uint64()
		return index < prevTotalElements
	}

	isEventLo := func(event *bindings.StateCommitmentChainStateBatchAppended, index uint64) bool {
		prevTotalElements := event.PrevTotalElements.Uint64()
		batchSize := event.BatchSize.Uint64()
		return index >= prevTotalElements+batchSize
	}

	lowerBound := uint64(0)
	upperBound := totalBatches.Uint64() - 1
	batchEvent, err := GetStateBatchAppendedEventByBatchIndex(mc, []*big.Int{big.NewInt(int64(upperBound))}, 0, 999)
	if err != nil {
		return nil, err
	}

	// Only happens when no batches have been submitted yet.
	if batchEvent == nil {
		return nil, nil
	}

	if isEventLo(batchEvent, transactionIndex) {
		// Upper bound is too low, means this transaction doesn't have a corresponding state batch yet.
		return nil, nil
	} else if !isEventHi(batchEvent, transactionIndex) {
		// Upper bound is not too low and also not too high. This means the upper bound event is the
		// one we're looking for! Return it.
		return batchEvent, nil
	}

	// Binary search to find the right event. The above checks will guarantee that the event does
	// exist and that we'll find it during this search.
	for lowerBound < upperBound {
		middleOfBounds := (lowerBound + upperBound) / 2
		batchEvent, err = GetStateBatchAppendedEventByBatchIndex(mc, []*big.Int{big.NewInt(int64(middleOfBounds))}, 0, 999)
		if err != nil {
			return nil, err
		}

		if isEventHi(batchEvent, transactionIndex) {
			upperBound = middleOfBounds
		} else if isEventLo(batchEvent, transactionIndex) {
			lowerBound = middleOfBounds
		} else {
			break
		}
	}

	return batchEvent, nil
}

func GetStateRootBatchByTransactionIndex(
	mc *MantleCenter,
	transactionIndex uint64,
) (srb StateRootBatch, err error) {
	stateBatchAppendedEvent, err := GetStateBatchAppendedEventByTransactionIndex(mc, transactionIndex)
	if err != nil {
		return srb, fmt.Errorf("GetStateBatchAppendedEventByTransactionIndex err: %v", err)
	}
	tx, _, err := mc.L1Client.TransactionByHash(context.Background(), stateBatchAppendedEvent.Raw.TxHash)
	if err != nil {
		return srb, fmt.Errorf("TransactionByHash err: %v", err)
	}

	abiJson := "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vmHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"inboxSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_batch\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_shouldStartAtElement\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"createAssertionWithStateBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
	_, inputArgs, err := lib.ParseContractFunctionInputs(abiJson, tx.Data())
	if err != nil {
		return srb, fmt.Errorf("ParseContractFunctionInputs err: %v", err)
	}

	for _, b := range inputArgs["_batch"].([][32]uint8) {
		srb.StateRoots = append(srb.StateRoots, fmt.Sprintf("0x%x", b))
	}

	srb.BlockNumber = stateBatchAppendedEvent.Raw.BlockNumber
	srb.Header.BatchIndex = stateBatchAppendedEvent.BatchIndex
	srb.Header.BatchRoot = stateBatchAppendedEvent.BatchRoot
	srb.Header.BatchSize = stateBatchAppendedEvent.BatchSize
	srb.Header.PrevTotalElements = stateBatchAppendedEvent.PrevTotalElements
	srb.Header.Signature = stateBatchAppendedEvent.Signature
	srb.Header.extraData = stateBatchAppendedEvent.ExtraData

	return srb, err
}

// func GetMessageStateRoot(mc *MantleCenter) error {

// }

func GetMessagesByTransaction(mc *MantleCenter, txHash common.Hash) (resolved Resolved, err error) {
	re, err := mc.L2Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return resolved, fmt.Errorf("TransactionReceipt err: %v", err)
	}

	abiJson := "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"SentMessage\",\"type\":\"event\"}]"
	contractAbi, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return resolved, fmt.Errorf("abi.JSON err: %v", err)
	}

	SentMessageSig := []byte("SentMessage(address,address,bytes,uint256,uint256)")
	SentMessageSigHash := crypto.Keccak256Hash(SentMessageSig)

	for _, vLog := range re.Logs {

		switch vLog.Topics[0].Hex() {
		case SentMessageSigHash.Hex():
			logdata, err := contractAbi.Unpack("SentMessage", vLog.Data)
			if err != nil {
				return resolved, fmt.Errorf("contractAbi.Unpackerr: %v", err)
			}

			resolved = Resolved{
				Direction:       1,
				Target:          mc.Env.AddressList.ProxyBVML1StandardBridge,
				Sender:          logdata[0].(common.Address),
				Message:         fmt.Sprintf("0x%x", logdata[1].([]uint8)),
				MessageNonce:    logdata[2].(*big.Int),
				Value:           0,
				MinGasLimit:     logdata[3].(*big.Int),
				LogIndex:        vLog.Index,
				BlockNumber:     vLog.BlockNumber,
				TransactionHash: vLog.TxHash,
			}
		}
	}

	return resolved, nil
}

func GetMessageStateRoot(mc *MantleCenter, txHash common.Hash) (sr StateRoot, err error) {
	re, err := mc.L2Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return sr, err
	}

	messageTxIndex := re.BlockNumber.Uint64() - 1
	stateRootBatch, err := GetStateRootBatchByTransactionIndex(mc, messageTxIndex)
	if err != nil {
		return sr, fmt.Errorf("GetStateRootBatchByTransactionIndex err: %v", err)
	}

	indexInBatch := messageTxIndex - stateRootBatch.Header.PrevTotalElements.Uint64()

	sr = StateRoot{
		StateRoot:             stateRootBatch.StateRoots[indexInBatch],
		StateRootIndexInBatch: indexInBatch,
		Batch:                 stateRootBatch,
	}
	return sr, nil
}

func GetMessageProof(mc *MantleCenter, txHash common.Hash) (map[string]interface{}, error) {
	resolved, err := GetMessagesByTransaction(mc, txHash)
	if err != nil {
		return nil, fmt.Errorf("GetMessagesByTransaction err: %v", err)
	}

	stateRoot, err := GetMessageStateRoot(mc, txHash)
	if err != nil {
		return nil, fmt.Errorf("GetMessageStateRoot err: %v", err)
	}

	abiJSON := `[{
		"inputs": [
		  {
			"internalType": "address",
			"name": "_target",
			"type": "address"
		  },
		  {
			"internalType": "address",
			"name": "_sender",
			"type": "address"
		  },
		  {
			"internalType": "bytes",
			"name": "_message",
			"type": "bytes"
		  },
		  {
			"internalType": "uint256",
			"name": "_messageNonce",
			"type": "uint256"
		  }
		],
		"name": "relayMessage",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	  }]`

	contractAbi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}

	data, err := contractAbi.Pack("relayMessage",
		common.HexToAddress(resolved.Target),
		resolved.Sender,
		common.Hex2Bytes(resolved.Message),
		resolved.MessageNonce)
	if err != nil {
		panic(err)
	}

	messengerAddress := remove0x(mc.Env.L2cdmAddress)

	msgHash1 := crypto.Keccak256([]byte(common.Bytes2Hex(data) + messengerAddress))
	msgHash2 := crypto.Keccak256(append(msgHash1, bytes.Repeat([]byte{0x00}, 32)...))

	messageSlot := hex.EncodeToString(msgHash2)
	_, err = MakeStateTrieProof(mc, resolved.BlockNumber, common.HexToAddress(mc.Env.BVM_L2ToL1MessagePasser), messageSlot)
	if err != nil {
		return nil, fmt.Errorf("makeStateTrieProof err: %v", err)
	}

	hexProof, err := MakeMerkleTreeProof(
		stateRoot.Batch.StateRoots,
		stateRoot.StateRootIndexInBatch,
	)
	if err != nil {
		return nil, fmt.Errorf("MakeMerkleTreeProof err: %v", err)
	}

	return map[string]interface{}{
		"stateRoot":            stateRoot.StateRoot,
		"stateRootBatchHeader": stateRoot.Batch.Header,
		"stateRootProof": map[string]interface{}{
			"index":    stateRoot.StateRootIndexInBatch,
			"siblings": hexProof,
		},
		// "stateTrieWitness":   toHexString(rlp.encode(stateTrieProof.accountProof)),
		// "storageTrieWitness": toHexString(rlp.encode(stateTrieProof.storageProof)),
	}, nil
}

func remove0x(str string) string {
	if len(str) >= 2 && str[:2] == "0x" {
		return str[2:]
	}
	return str
}

func MakeStateTrieProof(
	mc *MantleCenter,
	blockNumber uint64,
	address common.Address,
	slot string,
) (map[string]interface{}, error) {
	client, err := rpc.DialContext(context.Background(), mc.Env.L2URL)
	if err != nil {
		return nil, fmt.Errorf("rpc.DialContext err: %v", err)
	}

	var res Result
	err = client.CallContext(context.Background(),
		&res,
		"eth_getProof",
		address,
		[]string{slot},
		fmt.Sprintf("0x%x", blockNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("client.CallContext err: %v", err)
	}

	return map[string]interface{}{
		"accountProof": res.AccountProof,
		"storageProof": res.StorageProofs[0].Proof,
		"storageValue": res.StorageProofs[0].Value,
		"storageRoot":  res.StorageHash,
	}, nil
}

func MakeMerkleTreeProof(leaves []string, index uint64) ([]string, error) {

	fmt.Println(leaves)
	fmt.Println(index)

	correctedTreeSize := int(math.Pow(2, math.Ceil(math.Log2(float64(len(leaves))))))

	var parsedLeaves []string
	for i := 0; i < correctedTreeSize; i++ {
		if i < len(leaves) {
			parsedLeaves = append(parsedLeaves, leaves[i])
		} else {
			leaveTmp := keccak256(common.Hex2Bytes("0x" + string(bytes.Repeat([]byte{0}, 32))))
			parsedLeaves = append(parsedLeaves, "0x"+common.Bytes2Hex(leaveTmp))
		}
	}
	fmt.Println(parsedLeaves)
	// merkletreejs prefers things to be Buffers.
	var bufLeaves [][]byte
	for _, leaf := range parsedLeaves {
		bufLeaves = append(bufLeaves, common.FromHex(leaf))
	}

	tree := NewMerkleTree(bufLeaves)
	proof := tree.Proof(bufLeaves[4])

	for _, p := range proof {
		fmt.Println("proof: ", common.Bytes2Hex(p))
	}

	// proof, err := tree.GetProof(bufLeaves[index])
	// if err != nil {
	// 	return nil, err
	// }

	// var hexProof []string
	// for _, element := range proof {
	// 	hexProof = append(hexProof, common.Bytes2Hex(element.Data))
	// }

	// return hexProof, nil
	return nil, nil
}

func keccak256(data []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(data)
	return h.Sum(nil)
}

type Resolved struct {
	Direction       int64
	Target          string
	Sender          common.Address
	Message         string
	MessageNonce    *big.Int
	Value           int64
	MinGasLimit     *big.Int
	LogIndex        uint
	BlockNumber     uint64
	TransactionHash common.Hash
}

type StateRoot struct {
	StateRoot             string
	StateRootIndexInBatch uint64
	Batch                 StateRootBatch
}

type AccountProof struct {
	Proof []string `json:"proof"`
}

type StorageProof struct {
	Key   string   `json:"key"`
	Value string   `json:"value"`
	Proof []string `json:"proof"`
}

type Result struct {
	Address       string         `json:"address"`
	AccountProof  []string       `json:"accountProof"`
	Balance       string         `json:"balance"`
	CodeHash      string         `json:"codeHash"`
	Nonce         string         `json:"nonce"`
	StorageHash   string         `json:"storageHash"`
	StorageProofs []StorageProof `json:"storageProof"`
}

func DecodeAppendSequencerBatchParams(inputData string) (params AppendSequencerBatchParams, err error) {
	strs := strings.Split(inputData, "d0f89344")

	rawBytes, err := hex.DecodeString(strs[1])
	if err != nil {
		return params, fmt.Errorf("DecodeString err: %v", err)
	}

	err = params.Read(bytes.NewReader(rawBytes))
	if err != nil {
		return params, fmt.Errorf("params.Read err: %v", err)
	}

	return params, err
}

func FinalizeMessage(mc *MantleCenter, txHash common.Hash) {

	// command := fmt.Sprintf("node packages/finalize/index.js pri=%s l1url=%s l2url=%s hash=%s",
	// 	mc.Env.PrivateKeyList[0][0], mc.Env.L1URL, mc.Env.L2URL, txHash)
	cmd := exec.Command("bash", "-c", "npx ts-node packages/sdk/finalize.ts")
	addEnvList := []string{}
	addEnvList = append(addEnvList, fmt.Sprintf("L1_URL=%v", mc.Env.L1URL))
	addEnvList = append(addEnvList, fmt.Sprintf("L1_CHAINID=%v", mc.Env.L1ChainID))
	addEnvList = append(addEnvList, fmt.Sprintf("L2_URL=%v", mc.Env.L2URL))
	addEnvList = append(addEnvList, fmt.Sprintf("L2_CHAINID=%v", 1705003))
	addEnvList = append(addEnvList, fmt.Sprintf("PRIVATE_KEY=%v", mc.Env.PrivateKeyList[0][0]))
	addEnvList = append(addEnvList, fmt.Sprintf("ADDRESS_MANAGER_ADDRESS=%v", mc.Env.AddressList.AddressManager))
	addEnvList = append(addEnvList, fmt.Sprintf("L1_CROSS_DOMAIN_MESSENGER_ADDRESS=%v", mc.Env.AddressList.ProxyBVML1CrossDomainMessenger))
	addEnvList = append(addEnvList, fmt.Sprintf("L1_STANDARD_BRIDGE_ADDRESS=%v", mc.Env.AddressList.ProxyBVML1StandardBridge))
	addEnvList = append(addEnvList, fmt.Sprintf("STATE_COMMITMENT_CHAIN_ADDRESS=%v", mc.Env.AddressList.StateCommitmentChain))
	addEnvList = append(addEnvList, fmt.Sprintf("CANONICAL_TRANSACTION_CHAIN_ADDRESS=%v", mc.Env.AddressList.CanonicalTransactionChain))
	addEnvList = append(addEnvList, fmt.Sprintf("BOND_MANAGER_ADDRESS=%v", mc.Env.AddressList.BondManager))
	addEnvList = append(addEnvList, fmt.Sprintf("TestMantleToken_ADDRESS=%v", mc.Env.AddressList.Proxy__L1MantleToken))
	addEnvList = append(addEnvList, fmt.Sprintf("FRAUD_PROOF_WINDOW=%v", 60))
	addEnvList = append(addEnvList, fmt.Sprintf("TX_HASH=%v", txHash))
	cmd.Env = append(os.Environ(), addEnvList...)
	fmt.Println(addEnvList)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	return "", fmt.Errorf("error executing command: %v", err)
	// }
	// fmt.Println("===2")

	// return string(output), nil
	// 创建管道获取脚本的标准输出
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("创建标准输出管道出错：%v\n", err)
	}
	// 启动命令
	err = cmd.Start()
	if err != nil {
		fmt.Printf("启动命令出错：%v\n", err)
	}

	// 实时获取输出并处理
	scanner := bufio.NewScanner(stdoutPipe)
	for scanner.Scan() {
		line := scanner.Text()
		// 处理每行输出
		fmt.Println(line)
	}

	// 等待命令执行完成
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("等待命令执行完成出错：%v\n", err)
	}
}

func GetDataLogs(c *ethclient.Client, contractAddress common.Address, contractAbi abi.ABI, funcString string) error {

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("c.FilterLogs err: %v", err)
	}

	funcSig := []byte(funcString)
	funcHash := crypto.Keccak256Hash(funcSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case funcHash.Hex():
			logdata, err := contractAbi.Unpack(strings.Split(funcString, "(")[0], vLog.Data)
			if err != nil {
				return fmt.Errorf("contractAbi.Unpack err: %v", err)
			}
			fmt.Println(logdata)
		}
	}
	return nil
}

func GetMapLogs(c *ethclient.Client, contractAddress common.Address, contractAbi abi.ABI, funcString string, v map[string]interface{}) (
	blockNumber uint64,
	logIndex uint,
	err error,
) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := c.FilterLogs(context.Background(), query)
	if err != nil {
		return blockNumber, logIndex, fmt.Errorf("c.FilterLogs err: %v", err)
	}

	funcSig := []byte(funcString)
	funcHash := crypto.Keccak256Hash(funcSig)

	for _, vLog := range logs {
		blockNumber = vLog.BlockNumber
		logIndex = vLog.Index

		// for i, funcHash := range funcHashs {
		// 	if vLog.Topics[0].Hex() == funcHash.Hex() {
		// 		err := contractAbi.UnpackIntoMap(v, strings.Split(funcStrings[i], "(")[0], vLog.Data)
		// 		if err != nil {
		// 			return blockNumber, logIndex, fmt.Errorf("contractAbi.Unpack err: %v", err)
		// 		}
		// 	}
		// }
		switch vLog.Topics[0].Hex() {
		case funcHash.Hex():
			err := contractAbi.UnpackIntoMap(v, strings.Split(funcString, "(")[0], vLog.Data)
			if err != nil {
				return blockNumber, logIndex, fmt.Errorf("contractAbi.Unpack err: %v", err)
			}
		}
	}
	return blockNumber, logIndex, err
}

func SubscribeNewHead(c ethclient.Client) {
	// 监听新区块
	headers := make(chan *types.Header)
	sub, err := c.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Println("SubscribeNewHead err: ", err)
		return
	}

	for {
		select {
		case header := <-headers:
			fmt.Println("新区块号:", header.Number.String())
		case err := <-sub.Err():
			fmt.Println("sub err: ", err)
			return
		case <-time.After(10 * time.Second):
			// 等待新区块事件，可以根据需求调整等待时间
		}
	}
}

func SubscribeFilterLogs(c ethclient.Client, CA string) {

	contractAddress := common.HexToAddress(CA)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := c.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		fmt.Println("SubscribeFilterLogs err: ", err)
		return
	}

	for {
		select {
		case err := <-sub.Err():
			if err != nil {
				fmt.Println("SubscribeFilterLogs's for err: ", err)
				return
			}
		case vLog := <-logs:
			fmt.Println("===============start================")
			fmt.Println("vLog: ", vLog)
			fmt.Println("vLog.Address: ", vLog.Address)
			fmt.Println("vLog.Index: ", vLog.Index)
			fmt.Println("vLog.Data: ", vLog.Data)
			fmt.Println("vLog.Topics: ", vLog.Topics)
			fmt.Println("=================end==================")
		}
	}
}
