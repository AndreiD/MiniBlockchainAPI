package handlers

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/AndreiD/MiniBlockchainAPI/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/gin-gonic/gin"
)

// SendEth sends ETH
func SendEth(c *gin.Context) {

	toAddress := c.DefaultQuery("to_address", "")
	amountInWei := c.DefaultQuery("amount_in_wei", "")
	senderPrivateKey := c.DefaultQuery("sender_private_key", "")

	log.Printf("Sending %s wei to %s\n", amountInWei, toAddress)

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error reading private key"})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error getting publicKey"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := helpers.ETHClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	value := new(big.Int)
	value, ok = value.SetString(amountInWei, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong amount given"})
		return
	}

	gasLimit := uint64(55723) // in units

	gasPrice, err := helpers.ETHClient.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var data []byte //nil
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)

	chainID, err := helpers.ETHClient.NetworkID(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = helpers.ETHClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": value, "currency": "wei", "hash": signedTx.Hash().String()})

}

// SendToken sends tokens
func SendToken(c *gin.Context) {

	senderPrivateKey := c.DefaultQuery("sender_private_key", "")
	contract := c.DefaultQuery("contract", "")
	_toAddress := c.DefaultQuery("to_address", "")
	_amount := c.DefaultQuery("amount_in_wei", "")

	log.Printf("Sending %s tokens (contract: %s) to %s\n", _amount, contract, _toAddress)

	amount := new(big.Int)
	amount, ok := amount.SetString(_amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
		return
	}

	privateKey, err := crypto.HexToECDSA(senderPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {

		c.JSON(http.StatusBadRequest, gin.H{"error": "error casting public key to ECDSA"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := helpers.ETHClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	toAddress := common.HexToAddress(_toAddress)
	tokenAddress := common.HexToAddress(contract)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit := uint64(55723) // in units
	gasPrice, err := helpers.ETHClient.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	chainID, err := helpers.ETHClient.NetworkID(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = helpers.ETHClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": amount, "currency": "Tokens", "hash": signedTx.Hash().String()})

}
