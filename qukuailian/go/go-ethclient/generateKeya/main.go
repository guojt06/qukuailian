package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// 1. 生成一个新的ECDSA私钥（使用secp256k1椭圆曲线）
	// 这是以太坊账户的核心，必须绝对保密！
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 2. 将私钥转换为字节格式并十六进制编码显示
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// hexutil.Encode()会添加'0x'前缀，[2:]去掉前缀只显示纯十六进制字符串
	fmt.Println("私钥 (Private Key):", hexutil.Encode(privateKeyBytes)[2:])

	// 3. 从私钥推导出对应的公钥
	publicKey := privateKey.Public()

	// 4. 类型断言：确保公钥是ECDSA类型（secp256k1曲线）
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("无法断言类型: 公钥不是*ecdsa.PublicKey类型")
	}

	// 5. 将公钥转换为字节格式
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 显示原始公钥字节（去掉'0x04'前缀，这是非压缩公钥的标识字节）
	fmt.Println("原始公钥 (Raw Public Key):", hexutil.Encode(publicKeyBytes)[4:])

	// 6. 使用内置函数直接计算以太坊地址
	// PubkeyToAddress做了两件事：1) Keccak256哈希 2) 取后20字节
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("以太坊地址 (Ethereum Address):", address)

	// 7. 手动计算以太坊地址的过程（展示原理）
	hash := sha3.NewLegacyKeccak256() // 创建Keccak-256哈希器（注意：这是LegacyKeccak，不是标准SHA-3）

	// 8. 对公钥进行哈希（跳过第一个字节0x04，这是非压缩格式标识）
	hash.Write(publicKeyBytes[1:]) // 从第1个字节开始（跳过0x04）

	// 9. 显示完整的32字节哈希值
	fmt.Println("完整Keccak-256哈希 (Full Hash):", hexutil.Encode(hash.Sum(nil)[:]))

	// 10. 显示地址部分（哈希值的后20字节）
	// 这就是以太坊地址：Keccak256(公钥)的最后20个字节
	fmt.Println("地址部分 (Address part, last 20 bytes):", hexutil.Encode(hash.Sum(nil)[12:]))
}
