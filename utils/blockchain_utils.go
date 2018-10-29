package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

//GetBalance the balance in wei for a given address
func GetWeiBalance(address string, client *ethclient.Client) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// GetBlockNumber gets the block number
func GetBlockNumber(client *ethclient.Client) (string, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return header.Number.String(), nil
}

// IsSyncying returns false if it's not syncing
func IsSyncying(client *ethclient.Client) (bool, error) {
	sync, err := client.SyncProgress(context.Background())
	if err != nil {
		return false, err
	}
	if sync == nil {
		return false, nil
	}
	return true, nil
}

// GenerateWallet generates a new wallet
func GenerateWallet() (address string, privateKey string, err error) {
	pKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("could not generate the private key %s", err)
	}
	privateKeyBytes := crypto.FromECDSA(pKey)

	publicKey := pKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	privateKey = hexutil.Encode(privateKeyBytes)
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("Address: " + address)
	fmt.Println("Private Key: " + privateKey)

	return address, privateKey, nil
}

// GetGasCost calculates gas cost
func GetGasCost(client ethclient.Client, txHash string) *big.Int {
	receipt, _ := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	feesWei := new(big.Int)
	feesWei.Add(feesWei, new(big.Int).Mul(new(big.Int).SetUint64(receipt.GasUsed), gasPrice))
	fmt.Printf(">>>> Gas Paid  %d <<<<\n", feesWei)
	return feesWei
}

// QueryBlock sample query block and list transactions...
func QueryBlock(blockNumber int, client *ethclient.Client) error {
	number := big.NewInt(int64(blockNumber))
	block, err := client.BlockByNumber(context.Background(), number)
	if err != nil {
		return err
	}
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		//TODO: not working
		if msg, err := tx.AsMessage(types.HomesteadSigner{}); err != nil {
			fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
	}
	return nil
}

// SendEthToAddress send eth (in wei) to an address (from a hardcoded private key)
func SendEthToAddress(value *big.Int, client *ethclient.Client) (txHash string, e error) {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		return "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}

	gasLimit := uint64(21000) // in units

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", fmt.Errorf("can't get a gas price suggestion: %s", err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

// TokenBalance gets the token balance of an account
// func TokenBalance(from string, instance *contracts.MspToken) (*big.Int, error) {
// 	balance, err := instance.BalanceOf(nil, common.HexToAddress(from))
// 	if err != nil {
// 		return big.NewInt(-1), err
// 	}
// 	return balance, nil
// }

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

// ToDecimal wei to decimals
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

// CalcGasCost calculate gas cost given gas limit (units) and gas price (wei)
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}

// SigRSV signatures R S V returned as arrays
func SigRSV(isig interface{}) ([32]byte, [32]byte, uint8) {
	var sig []byte
	switch v := isig.(type) {
	case []byte:
		sig = v
	case string:
		sig, _ = hexutil.Decode(v)
	}

	sigstr := common.Bytes2Hex(sig)
	rS := sigstr[0:64]
	sS := sigstr[64:128]
	R := [32]byte{}
	S := [32]byte{}
	copy(R[:], common.FromHex(rS))
	copy(S[:], common.FromHex(sS))
	vStr := sigstr[128:130]
	vI, _ := strconv.Atoi(vStr)
	V := uint8(vI + 27)

	return R, S, V
}
