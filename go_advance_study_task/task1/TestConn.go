package task1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func TestConn() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("SEPOLIA_API_TOKEN")
	url := "https://sepolia.infura.io/v3/" + token
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(9839605)

	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header info:")
	fmt.Println(header.Number.Uint64())
	fmt.Println(header.Time)
	fmt.Println(header.Difficulty.Uint64())
	fmt.Println(header.Hash().Hex())
	fmt.Println()

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("block info:")
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(len(block.Transactions()))
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func TestSendETH() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	token := os.Getenv("SEPOLIA_API_TOKEN")
	url := "https://sepolia.infura.io/v3/" + token
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	privateToken := os.Getenv("SEPOLIA_ACCOUNT_1_TOKEN")
	privateKey, err := crypto.HexToECDSA(privateToken)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) //0.001
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x8732409019f906a40ee3c43e2bcc29f0c7855595")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signdTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signdTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signdTx.Hash().Hex())
}
