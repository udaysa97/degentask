package balancetest

import (
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

// InfuraURL is the Infura endpoint URL (replace with your own URL)
const InfuraURL = "https://mainnet.infura.io/v3/293d042d9e8744eb8d7da6e0d1937c5f"

// Transaction represents Ethereum transaction data
type Transaction struct {
	Hash        string `json:"hash"`
	BlockNumber uint64 `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Time        string `json:"time"`
}

// NFTTransfer represents NFT transfer event data
type NFTTransfer struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Token       string `json:"token"`
	Time        string `json:"time"`
	BlockNumber string
}

// ViewData represents the data to be rendered in the HTML template
type ViewData struct {
	WalletAddress string
	Transactions  []Transaction
	NFTTransfers  []NFTTransfer
}

func getTransactions(walletAddress, blockNumber string) ([]Transaction, error) {
	client, err := rpc.Dial(InfuraURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Get the latest block number if not specified

	// Retrieve transactions
	transactions := []Transaction{}
	err = client.Call(&transactions, "eth_getTransactions", walletAddress, blockNumber)
	if err != nil {
		return nil, err
	}

	// Convert timestamp to human-readable format
	for i, tx := range transactions {
		var block interface{}
		err := client.Call(&block, "eth_getBlockByNumber", tx.BlockNumber)
		if err != nil {
			return nil, err
		}
		//transactions[i].Time = time.Unix(int64(block.Timestamp), 0).UTC().Format(time.RFC3339)
		transactions[i].Time = time.Unix(int64(3), 0).UTC().Format(time.RFC3339)
	}

	return transactions, nil
}

func getNFTTransfers(walletAddress, blockNumber string) ([]NFTTransfer, error) {
	client, err := rpc.Dial(InfuraURL)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Get the latest block number if not specified
	// if blockNumber == "" {
	// 	blockNumber, err = client.EthBlockNumber()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// Retrieve NFT transfer events
	nftTransfers := []NFTTransfer{}
	err = client.Call(&nftTransfers, "eth_getNFTTransfers", walletAddress, blockNumber)
	if err != nil {
		return nil, err
	}

	// Convert timestamp to human-readable format
	// for i, transfer := range nftTransfers {
	// 	block, err := client.EthGetBlockByNumber(transfer.BlockNumber)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	nftTransfers[i].Time = time.Unix(int64(block.Timestamp), 0).UTC().Format(time.RFC3339)
	// }

	return nftTransfers, nil
}
