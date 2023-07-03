const ethers = require('ethers');
const mantleneSDK = require("@mantleio/sdk");
// const mantleneSDK = require("tianwei-qa-test");

function createSignerOrProvider(privateKey, providerUrl) {
    // 检查 privateKey 和 providerUrl 参数是否存在
    if (!privateKey || !providerUrl) {
      throw new Error('Private Key and Provider URL are required');
    }
  
    // 创建基于自定义 URL 的 Provider 对象
    const provider = new ethers.providers.JsonRpcProvider(providerUrl);
  
    // 创建钱包对象
    const wallet = new ethers.Wallet(privateKey, provider);

    // 返回根据参数创建的 Signer 或 Provider 对象
    return wallet;
}

// 获取命令行参数
const commandLineArgs = process.argv.slice(2);

// 解析命令行参数为键值对对象
const args = {};
commandLineArgs.forEach(arg => {
  const [key, value] = arg.split('=');
  args[key] = value;
});

let l1Wallet
let l2Wallet
let l1ChainId
let l2ChainId
let withdrawETHtx
let finalizeMessageResponse
let success = false

async function finalizeMessage(l1url, l2url, pri, txhash) {

    try {
        l1Wallet = createSignerOrProvider(pri, l1url)
      } catch (error) {
        console.error('l1Wallet Error:', error);
    }

    try {
        l2Wallet = createSignerOrProvider(pri, l2url)
      } catch (error) {
        console.error('l2Wallet Error:', error);
    }

    try {
      // 将 Chain ID 赋值给全局变量
      l1ChainId = (await l1Wallet.provider.getNetwork()).chainId;
    } catch (error) {
      console.error('l1ChainId Error:', error);
    }

    try {
        // 将 Chain ID 赋值给全局变量
        l2ChainId = (await l2Wallet.provider.getNetwork()).chainId;
    } catch (error) {
        console.error('l2ChainId Error:', error);
    }

    try {
        withdrawETHtx = await l2Wallet.provider.getTransaction(txhash)
    } catch (error) {
        console.error('getTransaction Error:', error);
    }

    let crossChainMessenger = new mantleneSDK.CrossChainMessenger({
        l1ChainId: l1ChainId,
        l2ChainId: l2ChainId,
        l1SignerOrProvider: l1Wallet,
        l2SignerOrProvider: l2Wallet,
    })

    while (!success) {
        try{
            finalizeMessageResponse = await crossChainMessenger.finalizeMessage(withdrawETHtx)
        } catch(error){
            console.log('getTransaction Error:', error);
            console.log("未完成 finalizeMessage……")
            await sleep(5000);
            continue
        }

        success = await isTransactionSuccessful(l1Wallet.provider, finalizeMessageResponse.hash);
        if (success){
            console.log("finalizeMessage hash: ", finalizeMessageResponse.hash)
            console.log("finalizeMessage type: ", finalizeMessageResponse.type)
            return true
        }else {
            throw new Error(`Failed to finalizeMessage receipt for txHash: ${finalizeMessageResponse.hash}`);
        }
    }
}

async function isTransactionSuccessful(provider, txHash) {

    while (true) {
        const receipt = await provider.getTransactionReceipt(txHash);
        if (receipt) {
            return receipt.status === 1;
        }
        await sleep(5000); // 等待1秒再继续查
    }
}

function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }
  

finalizeMessage(args["l1url"],args["l2url"],args["pri"],args["hash"]);
