package handlers

import (
	"math"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gitlab.com/AndreiDD/tokenominatorapi/bindings/cece"
	"gitlab.com/AndreiDD/tokenominatorapi/bindings/rp"
	"gitlab.com/AndreiDD/tokenominatorapi/bindings/viv"
	"gitlab.com/AndreiDD/tokenominatorapi/utils"
)

// GetETHPoaBalance gets the balance for a wallet in Ether (EV POA) (the thing that pays for the gas)
func GetETHPoaBalance(c *gin.Context) {

	address := c.DefaultQuery("address", "")

	balance, err := utils.GetWeiBalance(address, utils.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	c.JSON(http.StatusOK, gin.H{"balance": ethValue, "currency": "ETH (POA)"})
}

//GetETHRinkebyBalance ...
func GetETHRinkebyBalance(c *gin.Context) {

	address := c.DefaultQuery("address", "")

	balance, err := utils.GetWeiBalance(address, utils.RinkebyClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	c.JSON(http.StatusOK, gin.H{"balance": ethValue, "currency": "ETH (Rinkeby)"})
}

// GetCeCeBalance gets the balance for a token
func GetCeCeBalance(c *gin.Context) {

	contract := c.DefaultQuery("contract", "")
	address := c.DefaultQuery("address", "")

	instance, err := cece.NewCeCeToken(common.HexToAddress(contract), utils.RinkebyClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	c.JSON(http.StatusOK, gin.H{"balance": decimalValue, "currency": "CeCe Tokens"})
}

// GetRPBalance gets the balance for a token
func GetRPBalance(c *gin.Context) {

	contract := c.DefaultQuery("contract", "")
	address := c.DefaultQuery("address", "")

	instance, err := rp.NewRPToken(common.HexToAddress(contract), utils.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	c.JSON(http.StatusOK, gin.H{"balance": decimalValue, "currency": "RP Points"})
}

// GetVIVBalance gets the balance for a token
func GetVIVBalance(c *gin.Context) {

	contract := c.DefaultQuery("contract", "")
	address := c.DefaultQuery("address", "")

	instance, err := viv.NewVivToken(common.HexToAddress(contract), utils.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	decimalValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	c.JSON(http.StatusOK, gin.H{"balance": decimalValue, "currency": "VIV Points"})
}
