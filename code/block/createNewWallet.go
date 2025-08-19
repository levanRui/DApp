package task

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func CreateNewWallet() {
	// 生成新钱包的私钥
	privateKey, err := crypto.GenerateKey()
	fmt.Println("创建钱包私钥--------", privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//093f64cefef104e85968f8a6317a3ea0041352fa017733c9bb45514d09ee352b
	//已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥：
	//privateKey, err := crypto.HexToECDSA("093f64cefef104e85968f8a6317a3ea0041352fa017733c9bb45514d09ee352b")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//将私钥转为字节  go-ethereum hexutil 包将它转换为十六进制字符串，该包提供了一个带有字节切片的 Encode 方法。 然后我们在十六进制编码之后删除“0x”。
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("转化成私钥字节并去掉'0x'--------", hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'

	//获取公钥
	publicKey := privateKey.Public()
	fmt.Println("获取公钥publicKey-------", publicKey)
	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了 0x 和前 2 个字符 04，它是 EC 前缀，不是必需的。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fmt.Println("publicKeyECDSA---------", publicKeyECDSA)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("将公钥转为字节去掉0x04-------", hexutil.Encode(publicKeyBytes)[4:]) //去掉'0x04'
	//现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。 为了做到这一点，go-ethereum 加密包有一个 PubkeyToAddress 方法，
	//它接受一个 ECDSA 公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("生成公钥地址-------", address)
	//公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀。
	//以下是使用 golang.org/x/crypto/sha3 的 Keccak256 函数手动完成的方法。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("公钥的 Keccak-256 哈希---------", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("公钥的 Keccak-256 哈希最后的40个字符(20字节)--------", hexutil.Encode(hash.Sum(nil)[12:])) //原长32位，截去12位，保留后20位
}
