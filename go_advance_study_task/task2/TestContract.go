package task2

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/i159800/web3_homework/go_advance_study_task/task2/counter"
	"github.com/joho/godotenv"
)

func TestContract() {
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

	contractAddr := os.Getenv("CONTRACT_ADDR")
	counterContract, err := counter.NewCounter(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	privateToken := os.Getenv("SEPOLIA_ACCOUNT_1_TOKEN")
	privateKey, err := crypto.HexToECDSA(privateToken)
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := counterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := counterContract.Get(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value:", valueInContract)
}
