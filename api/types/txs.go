package types

type Transaction struct {
	Hash        string `json:"hash"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	BlockNumber uint64 `json:"blockNumber"`
	Time        string `json:"time"`
}

type BalanceData struct {
	WalletAddress string
	Transactions  []Transaction
}
