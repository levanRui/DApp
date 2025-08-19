package task

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func ETHTransfer() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/74285a1da0c344f88a9b4678681093ed")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA()
	if err != nil {
		log.Fatal(err)
	}
}
