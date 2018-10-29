package handlers

import (
	"math"
	"math/big"
	"net/http"

	"github.com/AndreiD/MiniBlockchainAPI/helpers"
	log "github.com/Sirupsen/logrus"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// GetEthBalance gets the balance for a wallet in Ether (the thing that pays for the gas)
func GetEthBalance(c *gin.Context) {

	address := c.DefaultQuery("address", "")

	balance, err := helpers.GetWeiBalance(address, helpers.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	log.Printf("Account %s has %s ETH\n", address, ethValue.String())
	c.JSON(http.StatusOK, gin.H{"balance": ethValue, "currency": "ETH"})
}

// GetTokenBalance gets the balance for a token
func GetTokenBalance(c *gin.Context) {

	contract := c.DefaultQuery("contract", "")
	address := c.DefaultQuery("address", "")

	tokenCaller, err := helpers.NewTokenCaller(common.HexToAddress(contract), helpers.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gets the balance
	bal, err := tokenCaller.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// gets the symbol
	symbol, err := tokenCaller.Symbol(nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	log.Printf("Account %s has %s %s\n", address, decimalValue.String(), symbol)
	c.JSON(http.StatusOK, gin.H{"balance": decimalValue, "currency": symbol})
}
