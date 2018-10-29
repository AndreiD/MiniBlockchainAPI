package handlers

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"

	"github.com/AndreiD/MiniBlockchainAPI/helpers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// SendEth sends ETH
func SendEth(c *gin.Context) {

	toAddress := c.DefaultQuery("to_address", "")

	// this is the private key for account 0x5d924b2d34643b4eb7d4291fdcb07236963f040f
	privateKey, err := crypto.HexToECDSA("908550C596A682C500FE1013EB3CEB5A8421FC62D6FF1F81CCDFEDD69768E560")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ">> error reading from private key"})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": ">> error getting publicKey"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := helpers.ETHClient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "where": ">> getting PendingNonceAt"})
		return
	}

	value := big.NewInt(100000000000000000) // in wei (0.1 eth)
	gasLimit := uint64(21000)               // in units
	gasPrice := big.NewInt(20000000)

	var data []byte //nil
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), value, gasLimit, gasPrice, data)

	chainID, err := helpers.ETHClient.NetworkID(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "where": ">> getting chainID"})
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "where": ">> SignTx"})
		return
	}

	err = helpers.ETHClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "where": ">> SendTransaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": value, "currency": "wei (POA)", "hash": signedTx.Hash().String()})

}

// SendToken sends tokens
func SendToken(c *gin.Context) {

	// this is the private key for sender's account 0x5d924b2d34643b4eb7d4291fdcb07236963f040f
	const senderPrivateKey = "908550C596A682C500FE1013EB3CEB5A8421FC62D6FF1F81CCDFEDD69768E560"
	const contractAddress = "0xabf59761226e415511ae828803cdf96142c31e89"

	toAddress := c.DefaultQuery("to_address", "")
	_amount := c.DefaultQuery("amount", "")
	amount := new(big.Int)
	amount, ok := amount.SetString(_amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
		return
	}

	hash, err := helpers.SignAndSendTokenTx(helpers.ETHClient, contractAddress, amount, toAddress, senderPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": amount, "currency": "VIV Tokens", "hash": hash})

}
