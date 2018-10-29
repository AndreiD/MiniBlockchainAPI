package utils

import (
	"fmt"

	log "github.com/Sirupsen/logrus"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// ETHClient the eth client
var ETHClient *ethclient.Client

// RPCClient the rpc client
var RPCClient *rpc.Client

// RinkebyClient the eth client
var RinkebyClient *ethclient.Client

// InitEthClient initializez the the client
func InitEthClient(url string) error {

	ethClient, err := ethclient.Dial(url)

	if err != nil {
		log.Errorf("Failed to connect to the ETH client: %v", err)
		return err
	}

	ETHClient = ethClient
	fmt.Printf("Connected to the ETH provider: %s\n", url)
	return nil

}

// InitRPCClient initializez the rpc client
func InitRPCClient(url string) error {

	rpcClient, err := rpc.Dial(url)

	if err != nil {
		log.Errorf("Failed to connect to the ETH client: %v", err)
		return err
	}
	fmt.Printf("Connected to the ETH RPC provider: %s\n", url)
	RPCClient = rpcClient
	return nil

}

// InitRinkeByClient - just hardcode it
func InitRinkeByClient() error {

	ethClient, err := ethclient.Dial("https://rinkeby.infura.io/2wIEW5lLyEaefGEPhoxX ")

	if err != nil {
		log.Errorf("Failed to connect to the ETH client: %v", err)
		return err
	}

	RinkebyClient = ethClient
	fmt.Println("Connected to the ETH provider: Rinkeby")
	return nil

}
