package userbalance

import (
	"context"
	"txcrawlerdegen/constants"
	blockchain "txcrawlerdegen/pkg/web3"
)

type UserBalanceSvc struct {
}

// If using builder pattern, init will be done with required constructor args, for sphinx demo skipping
func InitUserBalanceSvc() *UserBalanceSvc {
	return &UserBalanceSvc{}
}

func (*UserBalanceSvc) FetchUserBalance(ctx context.Context, walletAddress string, networkSymbol string) (string, error) {
	tokenAddress := constants.TOKEN_SYMBOL_ADDRESS_MAPPING[networkSymbol]
	// Return whatever is returned by called function
	return blockchain.FetchBalanceForToken(walletAddress, tokenAddress)

}

func (*UserBalanceSvc) FetchWalletTxsEth(ctx context.Context, walletAddress string) {

	// Return whatever is returned by called function
	blockchain.GetTransactions(walletAddress)

}
