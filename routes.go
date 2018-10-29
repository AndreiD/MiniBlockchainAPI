package main

import "gitlab.com/AndreiDD/tokenominatorapi/handlers"

// InitializeRoutes initializes all the routes in the app
func InitializeRoutes() {

	v1 := router.Group("/api/v1")
	{
		// info message
		v1.GET("/", handlers.Index)

		//shows health
		v1.GET("/health", handlers.Health)

		//================ BALANCES ================

		//shows the eth poa balance
		v1.GET("/balance/eth_poa", handlers.GetETHPoaBalance)

		//shows the eth poa balance
		v1.GET("/balance/eth_rinkeby", handlers.GetETHRinkebyBalance)

		//shows the CeCe token balance
		v1.GET("/balance/cece", handlers.GetCeCeBalance)

		//shows the RP token balance
		v1.GET("/balance/rp", handlers.GetRPBalance)

		//shows the VIV token balance
		v1.GET("/balance/viv", handlers.GetVIVBalance)

		//================ Transactions ================

		//send ETH POA to an address
		v1.PUT("/tx/eth_poa", handlers.SendEthPOA)

	}
}
