package task

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

/*
*
查询区块
*/
func QueryBlock() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/74285a1da0c344f88a9b4678681093ed")
	if err != nil {
		log.Fatal(err)
	}
	// 调用客户端的HeaderByNumber(context.Background(), blockNumber) 返回有关的一个区块的头信息
	blockNumber := big.NewInt(5671744)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	fmt.Println(header.Number.Uint64())     // 5671744
	fmt.Println(header.Time)                // 1712798400
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash().Hex())        // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	//调用客户端的BlockByNumber() 方法来获取完整区块
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number())     // 5671744  区块号
	fmt.Println(block.Time())       // 1712798400 区块时间戳
	fmt.Println(block.Difficulty()) // 0 区块难度

	fmt.Println(block.Hash().Hex())                                           // 区块hash
	fmt.Println(len(block.Transactions()))                                    // 70
	count, err := client.TransactionCount(context.Background(), block.Hash()) //区块交易数目
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count) //70
}
