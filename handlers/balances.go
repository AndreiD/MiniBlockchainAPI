package handlers

import (
	"math"
	"math/big"
	"net/http"

	"github.com/AndreiD/MiniBlockchainAPI/helpers"
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

	c.JSON(http.StatusOK, gin.H{"balance": ethValue, "currency": "ETH (POA)"})
}

// GetTokenBalance gets the balance for a token
func GetTokenBalance(c *gin.Context) {

	// contract := c.DefaultQuery("contract", "")
	// address := c.DefaultQuery("address", "")

	// instance, err := helpers.newTokenCaller(common.HexToAddress(contract), utils.RinkebyClient)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// fbalance := new(big.Float)
	// fbalance.SetString(bal.String())
	// decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	// c.JSON(http.StatusOK, gin.H{"balance": decimalValue, "currency": "CeCe Tokens"})
}
