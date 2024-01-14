package types

import (
	"fmt"
	"txcrawlerdegen/constants"
	"txcrawlerdegen/pkg/utils"
)

type FetchBalanceRequest struct {
	WalletAddress string `json:"wallet_address"`
	TokenSymbol   string `json:"token_symbol"`
}

type FetchBalanceResponse struct {
	Balance string `json:"balance"`
}

// Request validation
func (req *FetchBalanceRequest) Validate() error {
	if len(req.WalletAddress) <= 0 {
		return fmt.Errorf("Please enter a wallet address")
	}
	if !utils.IsValidBlockchainAddress(req.WalletAddress) {
		return fmt.Errorf("Please check wallet address and try again")
	}
	if len(req.TokenSymbol) <= 0 {
		req.TokenSymbol = "USDC"
	}
	if _, ok := constants.TOKEN_SYMBOL_ADDRESS_MAPPING[req.TokenSymbol]; !ok {
		validTokens := supportedTokensString()
		return fmt.Errorf("Please Send correct symbol, supported symbols:%s", validTokens)
	}
	return nil
}

// This function only exists since UI is not built. It returns suported tokens
// With UI this becomes a seperate API response
func supportedTokensString() string {
	tokensMap := constants.TOKEN_SYMBOL_ADDRESS_MAPPING
	supportedKeys := ""
	for key := range tokensMap {
		supportedKeys += key + ","
	}
	if len(supportedKeys) == 0 {
		return supportedKeys
	}
	return supportedKeys[:len(supportedKeys)-1]
}
