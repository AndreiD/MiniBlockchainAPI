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
  --header 'Content-Type: application/json'
~~~~


#### GET    /api/v1/balance/eth
gets the ETH balance.

params: address

Example:
~~~~
curl --request GET \
  --url 'http://localhost:9090/api/v1/balance/eth?address=0xF69D65f241a523837c7F37f8B38328176416E771' \
  --header 'Content-Type: application/json'
~~~~

#### GET    /api/v1/balance/token
gets the token balance.

params: contract, address

Example:
~~~~
curl --request GET \
  --url 'http://localhost:9090/api/v1/balance/token?contract=0x991c43f15b7d286f473e644df689dc3d722b58b2&address=0x5d924b2D34643B4Eb7D4291fDcb07236963f040f' \
  --header 'Content-Type: application/json'
~~~~

#### PUT    /api/v1/tx/send_eth
sends ETH to an address

params: to_address, sender_private_key, amount_in_wei

Example:
~~~~
curl --request PUT \
  --url 'http://localhost:9090/api/v1/tx/send_eth?to_address=0xF69D65f241a523837c7F37f8B38328176416E771&sender_private_key=908550C596A682C500FE1013EB3CEB5A8421FC62D6FF1F81CCDFEDD69768E560&amount_in_wei=100000000000000000' \
  --header 'Content-Type: application/json'
~~~~

#### PUT    /api/v1/tx/send_token
sends Tokens to an address

params: to_address, sender_private_key, contract, amount_in_wei

Example:
~~~~
curl --request PUT \
  --url 'http://localhost:9090/api/v1/tx/send_token?to_address=0xF69D65f241a523837c7F37f8B38328176416E771&sender_private_key=908550C596A682C500FE1013EB3CEB5A8421FC62D6FF1F81CCDFEDD69768E560&contract=0x991c43f15b7d286f473e644df689dc3d722b58b2&amount_in_wei=1000000000000000000' \
  --header 'Content-Type: application/json'
~~~~

#### License: GPL-3.0
