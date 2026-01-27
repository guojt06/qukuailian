package main

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
使用 Sepolia 测试网络实现基础的区块链交互，包括查询区块和发送交易。

	具体任务

环境搭建
安装必要的开发工具，如 Go 语言环境、 go-ethereum 库。
注册 Infura 账户，获取 Sepolia 测试网络的 API Key。
查询区块
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
实现查询指定区块号的区块信息，包括区块的哈希、时间戳、交易数量等。
输出查询结果到控制台。
发送交易
准备一个 Sepolia 测试网络的以太坊账户，并获取其私钥。
编写 Go 代码，使用 ethclient 连接到 Sepolia 测试网络。
构造一笔简单的以太币转账交易，指定发送方、接收方和转账金额。
对交易进行签名，并将签名后的交易发送到网络。
输出交易的哈希值。
*/

func main() {
	//// 连接到 Sepolia 测试网络
	//client, err := ethclient.Dial("https://sepolia.infura.io/v3/cf92a775b0ee4dfaaa0d4d845bfe4e9f")
	//if err != nil {
	//	log.Fatal("连接失败:", err)
	//}
	//defer client.Close()
	//
	//// 查询区块信息
	//block, err := client.BlockByNumber(context.Background(), big.NewInt(5671744))
	//if err != nil {
	//	log.Fatal("查询区块失败:", err)
	//}
	//
	//// 从私钥获取公钥和地址
	//_, err = crypto.HexToECDSA("")
	//if err != nil {
	//	log.Fatal("私钥解析失败:", err)
	//}
	//
	//// 输出区块信息
	//fmt.Println("=== 区块信息 ===")
	//fmt.Printf("区块号: %d\n", block.Number().Int64())
	//fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
	//fmt.Printf("区块时间戳: %d (Unix 时间)\n", block.Time())
	//fmt.Printf("交易数量: %d\n", len(block.Transactions()))
	//fmt.Printf("矿工地址: %s\n", block.Coinbase().Hex())
	//fmt.Printf("区块难度: %d\n", block.Difficulty().Uint64())
	//fmt.Printf("区块大小: %d 字节\n", block.Size())
	//fmt.Printf("父区块哈希: %s\n", block.ParentHash().Hex())
	//fmt.Printf("Gas 使用量: %d\n", block.GasUsed())
	//fmt.Printf("Gas 限制: %d\n", block.GasLimit())
	//fmt.Printf("区块随机数: %d\n", block.Nonce())
	//fmt.Println("===============")
	//
	//// 获取最新区块号
	//header, _ := client.HeaderByNumber(context.Background(), nil)
	//fmt.Printf("当前最新区块: %d\n", header.Number.Int64())

	txHash, _ := SendTransaction(0.0001)
	log.Fatal("交易已发送！哈希: %s\n", txHash)
	log.Fatal("查看交易: https://sepolia.etherscan.io/tx/%s\n", txHash)

}

// SendTransaction 发送以太币转账交易
func SendTransaction(amountInEther float64) (string, error) {
	// 连接到 Sepolia 测试网络
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cf92a775b0ee4dfaaa0d4d845bfe4e9f")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	// 从私钥获取公钥和地址
	privateKey, err := crypto.HexToECDSA("93f6a3ca82d0fa9b953b81ea5bb28d1f958b0c788b7bdecb4ba11f19c0301c61")
	if err != nil {
		log.Fatal("私钥解析失败:", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥转换失败:", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取 nonce 失败:", err)
	}

	// 获取当前 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取 gas 价格失败:", err)
	}

	// 设置接收地址
	toAddress := common.HexToAddress("0x1f25e88BAa52089072e1d1025A5Baf578cB5A837")

	// 转换转账金额：以太币 → Wei
	// 1 ETH = 10^18 Wei
	amount := big.NewInt(int64(amountInEther * 1e18))

	// 获取链 ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取链 ID 失败:", err)
	}

	// 创建交易
	tx := types.NewTransaction(
		nonce,
		toAddress,
		amount,
		uint64(21000), // 标准转账的 gas 限制
		gasPrice,
		nil,
	)

	// 签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("交易签名失败:", err)
	}

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("发送交易失败:", err)
	}

	log.Fatal("交易已发送！哈希: %s\n", signedTx.Hash().Hex())

	// 返回交易哈希
	return signedTx.Hash().Hex(), nil
}
