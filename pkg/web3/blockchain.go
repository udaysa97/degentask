package blockchain

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
	"txcrawlerdegen/internal/appconfig"
	"txcrawlerdegen/pkg/logger"

	tt "github.com/0fatih/ethereum-with-go/erc20Token/token"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraclient *ethclient.Client

const tokenABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"type":"function"}]`

func fetchInfuraClient() *ethclient.Client {
	var err error
	if infuraclient == nil {
		// Connecting with infura
		infuraclient, err = ethclient.Dial(appconfig.NODE_PROVIDER)
		if err != nil {
			logger.Error("Could not create connection with infura")
		}
	}
	return infuraclient
}

func FetchBalanceForToken(walletAddress string, token string) (string, error) {
	ethclient := fetchInfuraClient()
	tokenAddress := common.HexToAddress(token)
	// Using 0faith lib to avoid adding ABIs. Cleaner abstract solution for demo
	instance, err := tt.NewToken(tokenAddress, ethclient)
	if err != nil {
		logger.Error(err.Error())
		return "", fmt.Errorf("Internal Error")
	}

	address := common.HexToAddress(walletAddress)
	// Querying balance here. If required can query decimals etc. Not required for demo
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		logger.Error(err.Error())
		return "", fmt.Errorf(fmt.Sprintf("Could not fetch balance:%s", err.Error()))
	}
	// Converting to decimal(base10)
	return bal.Text(10), nil

	// tokenContract := bind.NewBoundContract(tokenAddress, parsedABI, ethclient, ethclient, ethclient)
	// balance := new(big.Int)
	// callOpts := &bind.CallOpts{
	// 	From:    account,
	// 	Pending: false,
	// }
	// var result *[]interface{}
	// err = tokenContract.Call(callOpts, result, "balanceOf", account)
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve token balance: %v", err)
	// }

}

type BlockInfo struct {
	Timestamp time.Time `json:"timestamp"`
}

// func getTransactionDetails(hash string, client *rpc.Client) (types.Transaction, error) {
// 	tx := types.Transaction{}
// 	err := client.Call(&tx, "eth_getTransactionByHash", hash)
// 	if err != nil {
// 		return types.Transaction{}, err
// 	}

// 	// Convert timestamp to human-readable format
// 	info := BlockInfo{}
// 	err = client.Call(&info, "eth_getBlockByNumber", tx.BlockNumber)
// 	if err != nil {
// 		return types.Transaction{}, err
// 	}
// 	tx.Time = info.Timestamp.String()

// 	return tx, nil
// }

func GetTransactions(walletAddress string) {
	client, err := ethclient.Dial(appconfig.NODE_PROVIDER)
	if err != nil {
		log.Fatal(err)
	}

	// Specify the wallet address
	address := common.HexToAddress("0xD1d6bF74282782B0b3eb1413c901D6eCF02e8e28")

	// Get the latest block number
	// header, err := client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Filter the transactions
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(35240892),
		ToBlock:   big.NewInt(35240893),
		Addresses: []common.Address{address},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("transactions.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file.WriteString("<html><body><table border=\"1\"><tr><th>Block Number</th><th>Transaction Hash</th><th>From</th><th>To</th><th>Value</th></tr>\n")

	// Iterate over the transactions
	for _, vLog := range logs {
		// Decode the transaction
		tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			log.Fatal(err)
		}
		logger.Debug(tx.ChainId().String())

		// Write the transaction details to the HTML file
		//file.WriteString("<tr><td>" + strconv.Itoa(int(tx.Nonce().Uint64())) + "</td><td>" + tx.Hash().Hex() + "</td><td>" + tx.From().Hex() + "</td><td>" + tx.To().Hex() + "</td><td>" + strings.TrimRight(tx.Value().String(), "\x00") + "</td></tr>\n")
	}

	// Write the HTML footer
	file.WriteString("</table></body></html>")
}

type EtherscanResponse struct {
	Result []struct {
		Hash            string `json:"hash"`
		Nonce           string `json:"nonce"`
		BlockNumber     string `json:"blockNumber"`
		TimeStamp       string `json:"timeStamp"`
		From            string `json:"from"`
		To              string `json:"to"`
		Value           string `json:"value"`
		ContractAddress string `json:"contractAddress"`
		Input           string `json:"input"`
		Type            string `json:"type"`
	} `json:"result"`
}

func ListNfts(walletAddress string) {
	resp, err := http.Get(fmt.Sprintf("https://api.etherscan.io/api?module=account&action=tokentx&address=%s&startblock=0&sort=asc&apikey=%s", walletAddress, appconfig.ETHERSCAN_TOKEN))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var etherscanResp EtherscanResponse
	json.Unmarshal(body, &etherscanResp)

	for _, tx := range etherscanResp.Result {
		fmt.Printf("Hash: %s, From: %s, To: %s, Contract Address: %s\n", tx.Hash, tx.From, tx.To, tx.ContractAddress)
	}
}
