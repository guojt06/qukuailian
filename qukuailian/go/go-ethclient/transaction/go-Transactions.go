package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 连接到以太坊 Sepolia 测试网络的公共节点
	// 注意：此公共节点可能不稳定，如果连接失败，可尝试更换为其他节点，如 "https://rpc.sepolia.org"
	client, err := ethclient.Dial("https://ethereum-sepolia.publicnode.com")
	if err != nil {
		log.Fatal(err) // 连接失败则终止程序
	}

	// 2. 获取当前网络的链 ID (Chain ID)
	// Chain ID 用于区分不同的以太坊网络（如主网为1，Sepolia为11155111），对交易签名至关重要
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 3. 指定一个具体的区块号，然后获取该区块的完整信息
	blockNumber := big.NewInt(5671744) // 创建一个代表区块 5671744 的大整数
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err) // 获取区块失败则终止
	}

	// 4. 遍历该区块中的所有交易，并打印第一条交易的详细信息
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 打印交易哈希 (交易的唯一ID)
		fmt.Println(tx.Value().String())    // 打印交易转移的以太币金额 (以Wei为单位)
		fmt.Println(tx.Gas())               // 打印该交易设置的Gas上限
		fmt.Println(tx.GasPrice().Uint64()) // 打印该交易愿意支付的每单位Gas价格
		fmt.Println(tx.Nonce())             // 打印发送者账户的交易序列号，用于防止重放攻击
		fmt.Println(tx.Data())              // 打印交易的附加数据 (调用合约时的输入数据)
		fmt.Println(tx.To().Hex())          // 打印交易接收者的地址

		// 5. 从交易签名中恢复出发送者的以太坊地址
		// 使用 EIP-155 签名器并传入 Chain ID 可以安全地恢复出签名者地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 打印发送者地址
		} else {
			log.Fatal(err) // 恢复发送者失败 (例如签名无效)
		}

		// 6. 根据交易哈希获取该交易的“收据”
		// 收据包含了交易执行后的结果信息，如是否成功、消耗的Gas、触发的日志等
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 打印交易状态: 1 表示成功，0 表示失败
		fmt.Println(receipt.Logs)   // 打印交易执行过程中产生的日志 (智能合约事件会存放在这里)
		break                       // 只处理第一条交易后就跳出循环
	}

	// 7. 根据区块哈希，获取该区块中的交易总数 (另一种方法)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	// 8. 遍历该区块中的所有交易，通过“区块哈希 + 交易索引”的方式获取每笔交易
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex()) // 打印通过此方法获取到的交易哈希
		break                        // 同样只处理第一条后跳出
	}

	// 9. 直接根据交易哈希获取一笔特定的交易
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	// 注意：这里有两个 fmt.Println，是原代码的笔误，第二个是完整的
	fmt.Println(isPending)       // 打印该交易是否还在待处理 (内存池) 中，false表示已上链
	fmt.Println(tx.Hash().Hex()) // 再次打印这笔交易的哈希
	// .Println(isPending)       // 这行是多余的，会编译错误，应删除
}
