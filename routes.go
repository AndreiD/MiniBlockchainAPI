package main

import "github.com/AndreiD/MiniBlockchainAPI/handlers"

// InitializeRoutes initializes all the routes in the app
func InitializeRoutes() {

	v1 := router.Group("/api/v1")
	{
		// info message
		v1.GET("/", handlers.Index)

		//shows health
		v1.GET("/health", handlers.Health)

		//================ BALANCES ================

		//shows the eth balance
		v1.GET("/balance/eth", handlers.GetEthBalance)

		//shows the CeCe token balance
		v1.GET("/balance/token", handlers.GetTokenBalance)

		//================ Transactions ================

		//send ETH POA to an address
		v1.PUT("/tx/send_eth", handlers.SendEth)

		//send Token to an address
		v1.PUT("/tx/send_token", handlers.SendToken)

	}
}
