package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 连接到以太坊的Sepolia测试网络节点
	//    使用以太坊基金会提供的公共免费节点
	client, err := ethclient.Dial("https://ethereum-sepolia.publicnode.com")
	if err != nil {
		log.Fatal(err) // 如果连接失败，记录错误并终止程序
	}

	// 2. 设置要查询的区块号
	//    这里指定查询第5671744号区块
	blockNumber := big.NewInt(5671744)

	// 3. 查询指定区块的区块头信息
	//    区块头包含区块的元数据，但不包含具体的交易详情
	header, err := client.HeaderByNumber(context.Background(), blockNumber)

	// 4. 打印区块头中的信息（先打印，后检查错误，注意这种写法不推荐用于生产环境）
	fmt.Println(header.Number.Uint64())     // 区块高度：5671744
	fmt.Println(header.Time)                // 区块生成时间戳（Unix时间）
	fmt.Println(header.Difficulty.Uint64()) // 区块挖矿难度（Sepolia是PoS链，所以难度为0）
	fmt.Println(header.Hash().Hex())        // 区块的哈希值

	// 5. 检查区块头查询是否出错
	if err != nil {
		log.Fatal(err) // 如果查询出错，记录错误并终止程序
	}

	// 6. 查询完整的区块信息（包含区块头和所有交易数据）
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 如果查询出错，记录错误并终止程序
	}

	// 7. 打印完整区块中的信息
	fmt.Println(block.Number().Uint64())     // 区块高度：与header.Number相同
	fmt.Println(block.Time())                // 区块时间戳：与header.Time相同
	fmt.Println(block.Difficulty().Uint64()) // 区块难度：与header.Difficulty相同
	fmt.Println(block.Hash().Hex())          // 区块哈希：与header.Hash相同
	fmt.Println(len(block.Transactions()))   // 该区块包含的交易数量

	// 8. 另一种获取区块交易数量的方法：通过区块哈希查询交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err) // 如果查询出错，记录错误并终止程序
	}

	fmt.Println(count) // 交易数量：与len(block.Transactions())相同，应为70
}
