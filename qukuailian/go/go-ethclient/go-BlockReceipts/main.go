package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// 1. 连接到以太坊Sepolia测试网络
	// 使用Alchemy节点服务，需要将<API_KEY>替换为你在Alchemy平台申请的实际API密钥
	client, err := ethclient.Dial("https://ethereum-sepolia.publicnode.com")
	if err != nil {
		log.Fatal(err) // 连接失败，终止程序
	}

	// 2. 定义要查询的区块信息
	blockNumber := big.NewInt(5671744) // 指定区块号
	// 指定区块哈希（与blockNumber对应的是同一个区块）
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")

	// 3. 通过区块哈希获取该区块内所有交易的回执
	// BlockReceipts方法可以一次获取一个区块内所有交易的回执，比逐个获取更高效
	// rpc.BlockNumberOrHashWithHash 将区块哈希包装成可接受的参数格式
	// 第二个参数false表示不要求返回完整的对象（只返回哈希）
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err) // 获取失败，终止程序
	}

	// 4. 通过区块号获取同一区块内所有交易的回执
	// rpc.BlockNumberOrHashWithNumber 将区块号包装成可接受的参数格式
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err) // 获取失败，终止程序
	}

	// 5. 验证两种方式获取的回执是否相同
	// 比较通过哈希和通过区块号获取的第一个交易回执是否为同一个对象（内存地址比较）
	// 注意：这里比较的是指针，true表示两个变量引用的是内存中的同一个回执对象
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // 应该输出 true

	// 6. 遍历并打印通过哈希获取的回执信息（只处理第一个）
	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)           // 交易状态：1表示成功，0表示失败
		fmt.Println(receipt.Logs)             // 交易产生的日志数组，这里是空数组
		fmt.Println(receipt.TxHash.Hex())     // 交易哈希的十六进制字符串
		fmt.Println(receipt.TransactionIndex) // 交易在区块中的索引位置（从0开始）
		// 如果是合约创建交易，这里会显示合约地址；否则显示零地址
		fmt.Println(receipt.ContractAddress.Hex()) // 合约地址（零地址表示这不是合约创建交易）
		break                                      // 只处理第一个回执就退出循环
	}

	// 7. 通过交易哈希单独获取交易回执
	// 这是获取单个交易回执的常用方法，适合已知具体交易哈希的场景
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err) // 获取失败，终止程序
	}

	// 8. 打印单独获取的交易回执信息
	fmt.Println(receipt.Status)                // 交易状态
	fmt.Println(receipt.Logs)                  // 交易日志
	fmt.Println(receipt.TxHash.Hex())          // 交易哈希
	fmt.Println(receipt.TransactionIndex)      // 交易在区块中的索引
	fmt.Println(receipt.ContractAddress.Hex()) // 合约地址
}
