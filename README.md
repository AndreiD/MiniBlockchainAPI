~~~~
███╗   ███╗██╗███╗   ██╗██╗    ██████╗ ██╗      ██████╗  ██████╗██╗  ██╗ ██████╗██╗  ██╗ █████╗ ██╗███╗   ██╗
████╗ ████║██║████╗  ██║██║    ██╔══██╗██║     ██╔═══██╗██╔════╝██║ ██╔╝██╔════╝██║  ██║██╔══██╗██║████╗  ██║
██╔████╔██║██║██╔██╗ ██║██║    ██████╔╝██║     ██║   ██║██║     █████╔╝ ██║     ███████║███████║██║██╔██╗ ██║
██║╚██╔╝██║██║██║╚██╗██║██║    ██╔══██╗██║     ██║   ██║██║     ██╔═██╗ ██║     ██╔══██║██╔══██║██║██║╚██╗██║
██║ ╚═╝ ██║██║██║ ╚████║██║    ██████╔╝███████╗╚██████╔╝╚██████╗██║  ██╗╚██████╗██║  ██║██║  ██║██║██║ ╚████║
╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝╚═╝    ╚═════╝ ╚══════╝ ╚═════╝  ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝  ╚═══╝
~~~~

## Send ETH & Tokens API

This small backend enables you to send ETH & Tokens to Ethereum based blockchains.

## API's

#### GET    /api/v1/
general info about the connection with the blockchain

#### GET    /api/v1/health
health about how much resources this mini-server takes

params: address

Example:
~~~~
curl --request GET \
  --url 'http://localhost:9090/api/v1/balance/token?contract=0x991c43f15b7d286f473e644df689dc3d722b58b2&address=0x5d924b2D34643B4Eb7D4291fDcb07236963f040f' \
  --header 'Authorization: Token abc' \
  --header 'Content-Type: application/json'
~~~~


#### GET    /api/v1/balance/eth
gets the ETH balance.

#### GET    /api/v1/balance/token
gets the token balance.

#### PUT    /api/v1/tx/send_eth
sends ETH to an address

#### PUT    /api/v1/tx/send_token
sends Tokens to an address

#### License: GPL-3.0
