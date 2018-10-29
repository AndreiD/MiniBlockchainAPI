package handlers

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"gitlab.com/AndreiDD/tokenominatorapi/utils"
)

// Index shows an info message
func Index(c *gin.Context) {

	blockNumber, err := utils.GetBlockNumber(utils.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	syncing, err := utils.IsSyncying(utils.ETHClient)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "alive", "block_number": blockNumber, "syncing": syncing})
}

// Health Endpoint
func Health(c *gin.Context) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.JSON(http.StatusOK, gin.H{"alloc": bToMb(m.Alloc), "total_alloc": bToMb(m.TotalAlloc),
		"sys": bToMb(m.Sys), "num_gc": m.NumGC})
}

// converts bytes to Mb
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
