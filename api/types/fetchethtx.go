package types

import (
	"fmt"
	"txcrawlerdegen/pkg/utils"
)

type FetchEthTxRequest struct {
	WalletAddress string `json:"wallet_address"`
	TokenSymbol   string `json:"token_symbol"`
}

type FetchEthTxResponse struct {
	Balance string `json:"balance"`
}

// Request validation
func (req *FetchEthTxRequest) Validate() error {
	if len(req.WalletAddress) <= 0 {
		return fmt.Errorf("Please enter a wallet address")
	}
	if !utils.IsValidBlockchainAddress(req.WalletAddress) {
		return fmt.Errorf("Please check wallet address and try again")
	}
	return nil
}
