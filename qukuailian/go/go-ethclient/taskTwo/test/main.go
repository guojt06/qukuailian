package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"myproject/go-ethclient/taskTwo/test/store"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 1. 连接到以太坊节点
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/cf92a775b0ee4dfaaa0d4d845bfe4e9f")
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer client.Close()

	// 2. 加载私钥
	privateKey, err := crypto.HexToECDSA("93f6a3ca82d0fa9b953b81ea5bb28d1f958b0c788b7bdecb4ba11f19c0301c61")
	if err != nil {
		log.Fatal("私钥加载失败:", err)
	}

	// 3. 准备部署参数
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥类型错误")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前区块号，验证连接
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("获取区块号失败:", err)
	}
	fmt.Printf("📡 已连接到Sepolia网络，当前区块: %d\n", blockNumber)

	// 4. 部署合约
	fmt.Println("🚀 开始部署合约...")

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("获取nonce失败:", err)
	}
	fmt.Printf("📝 账户nonce: %d\n", nonce)

	// 获取gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("获取gas价格失败:", err)
	}
	fmt.Printf("⛽ 建议gas价格: %d wei\n", gasPrice)

	// 获取链ID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取链ID失败:", err)
	}
	fmt.Printf("🆔 链ID: %d\n", chainId)

	// 创建交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal("创建交易签名器失败:", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	// 部署合约
	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal("部署交易发送失败:", err)
	}

	fmt.Println("\n📤 部署交易已发送到网络")
	fmt.Println("📝 合约地址:", address.Hex())
	fmt.Println("🔗 交易哈希:", tx.Hash().Hex())
	fmt.Println("🌐 在Etherscan查看: https://sepolia.etherscan.io/tx/" + tx.Hash().Hex())

	// ⭐⭐⭐ 关键：等待部署确认 ⭐⭐⭐
	fmt.Println("\n⏳ 等待部署确认（Sepolia网络约12-15秒/区块）...")

	if err := waitForDeployment(client, tx, address); err != nil {
		log.Fatal("部署失败:", err)
	}

	// ⭐ 现在可以安全调用合约了 ⭐
	fmt.Println("\n✅ 合约已成功部署！开始调用合约函数...")

	// 调用版本函数
	if err := callVersionFunctions(client, instance, address); err != nil {
		log.Fatal("调用版本函数失败:", err)
	}

	// 调用setItem函数
	if err := callSetItemFunction(client, privateKey, instance, address); err != nil {
		log.Fatal("调用setItem失败:", err)
	}

	fmt.Println("\n🎉 所有操作完成！")
}

// ⭐ 关键函数：等待合约部署完成
func waitForDeployment(client *ethclient.Client, tx *types.Transaction, contractAddress common.Address) error {
	ctx := context.Background()

	// 方法1：使用WaitMined等待交易被打包
	fmt.Println("1. 等待交易被打包进区块...")
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待交易失败: %v", err)
	}

	// 检查交易状态
	if receipt.Status == 0 {
		return fmt.Errorf("交易执行失败！请在Etherscan查看详情")
	}

	fmt.Printf("   ✅ 交易已确认！区块: %d, Gas使用: %d\n", receipt.BlockNumber, receipt.GasUsed)

	// 方法2：验证合约代码确实存在
	fmt.Println("2. 验证合约代码...")

	// 有时需要多等一会儿，合约代码才会完全可用
	for i := 1; i <= 10; i++ {
		code, err := client.CodeAt(ctx, contractAddress, nil)
		if err != nil {
			fmt.Printf("   尝试 %d/10: 查询代码失败: %v\n", i, err)
		} else if len(code) > 0 {
			fmt.Printf("   ✅ 合约代码已部署！长度: %d 字节\n", len(code))
			return nil
		} else {
			fmt.Printf("   尝试 %d/10: 代码长度: 0 (合约可能还在初始化)\n", i)
		}

		if i < 10 {
			time.Sleep(5 * time.Second) // 等待5秒再试
		}
	}

	return fmt.Errorf("等待合约代码超时")
}

// 调用版本相关函数
func callVersionFunctions(client *ethclient.Client, instance *store.Store, contractAddress common.Address) error {
	fmt.Println("\n📋 调用版本函数:")

	// 首先确认合约有代码
	code, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		return fmt.Errorf("查询合约代码失败: %v", err)
	}
	if len(code) == 0 {
		return fmt.Errorf("合约代码不存在")
	}
	fmt.Printf("   合约代码长度: %d 字节\n", len(code))

	// 调用自动生成的Version()函数
	fmt.Println("   1. 调用Version()...")
	version, err := instance.Version(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("调用Version()失败: %v", err)
	}
	fmt.Printf("      ✅ 版本: %s\n", version)

	// 调用自定义的GetVersion()函数
	fmt.Println("   2. 调用GetVersion()...")
	version2, err := instance.GetVersion(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("调用GetVersion()失败: %v", err)
	}
	fmt.Printf("      ✅ GetVersion(): %s\n", version2)

	// 调用GetVersionWithBlock()
	fmt.Println("   3. 调用GetVersionWithBlock()...")
	version3, blockNum, err := instance.GetVersionWithBlock(&bind.CallOpts{})
	if err != nil {
		return fmt.Errorf("调用GetVersionWithBlock()失败: %v", err)
	}
	fmt.Printf("      ✅ 版本: %s, 区块号: %d\n", version3, blockNum)

	return nil
}

// 调用setItem函数
func callSetItemFunction(client *ethclient.Client, privateKey *ecdsa.PrivateKey, instance *store.Store, contractAddress common.Address) error {
	fmt.Println("\n🔄 准备调用setItem函数...")

	// 获取新的nonce
	fromAddress := crypto.PubkeyToAddress(*privateKey.Public().(*ecdsa.PublicKey))
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("获取nonce失败: %v", err)
	}
	fmt.Printf("   账户新nonce: %d\n", nonce)

	// 获取新的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("获取gas价格失败: %v", err)
	}

	// 获取链ID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("获取链ID失败: %v", err)
	}

	// 创建新的交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return fmt.Errorf("创建交易签名器失败: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(100000)
	auth.GasPrice = gasPrice

	// 准备参数
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("myTestKey"))
	copy(value[:], []byte("myTestValue"))

	// 调用setItem
	fmt.Println("   调用setItem...")
	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		return fmt.Errorf("调用setItem失败: %v", err)
	}

	fmt.Printf("   ✅ setItem交易已发送！哈希: %s\n", tx.Hash().Hex())
	fmt.Println("   🌐 查看交易: https://sepolia.etherscan.io/tx/" + tx.Hash().Hex())

	// 等待setItem交易确认
	fmt.Println("   等待setItem交易确认...")
	ctx := context.Background()
	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		return fmt.Errorf("等待setItem交易失败: %v", err)
	}

	if receipt.Status == 0 {
		return fmt.Errorf("setItem交易执行失败")
	}

	fmt.Printf("   ✅ setItem交易成功！区块: %d\n", receipt.BlockNumber)

	return nil
}
