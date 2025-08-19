package task

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

func QueryReceipt() {
	//ethclient.Dial 是 Go 语言中以太坊客户端库 go-ethereum 里用于创建与以太坊节点连接的函数
	//rawurl：这是一个字符串类型的参数，代表的是以太坊节点的 RPC 接口地址。可以是 HTTP、HTTPS、WebSocket 等不同协议的地址
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/74285a1da0c344f88a9b4678681093ed")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	//blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	//common.HexToHash 是 Go 语言中以太坊客户端库 go-ethereum 里的一个实用函数，其主要作用是将十六进制字符串转换为以太坊的哈希类型 common.Hash
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	//通过区块哈希精确获取某个区块的信息（如区块头、交易列表）
	receiptsByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}
	//通过区块号获取某个区块的信息（如区块头、交易列表）
	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Uint64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptsByHash[0] == receiptsByNum[0]) //true
	//for _, receipt := range receiptsByHash {
	//	fmt.Println(receipt.Status)
	//	fmt.Println(receipt.Logs)
	//	fmt.Println(receipt.TxHash.Hex())
	//	fmt.Println(receipt.TransactionIndex)
	//	fmt.Println(receipt.ContractAddress)
	//}
	// 2.使用交易哈希查询收据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)
	fmt.Println(receipt.Logs)
	fmt.Println(receipt.TxHash.Hex())
	fmt.Println(receipt.TransactionIndex)
	fmt.Println(receipt.ContractAddress)
}
