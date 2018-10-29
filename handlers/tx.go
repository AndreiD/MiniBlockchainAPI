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

	gasLimit := uint64(21000) // in units

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
	toAddress := c.DefaultQuery("to_address", "")
	_amount := c.DefaultQuery("amount_in_wei", "")

	log.Printf("Sending %s tokens (contract: %s) to %s\n", _amount, contract, toAddress)

	amount := new(big.Int)
	amount, ok := amount.SetString(_amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
		return
	}

	hash, err := helpers.SignAndSendTokenTx(helpers.ETHClient, contract, amount, toAddress, senderPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": amount, "currency": "Tokens", "hash": hash})

}
