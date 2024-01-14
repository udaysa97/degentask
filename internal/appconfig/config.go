package appconfig

import (
	"os"
	"txcrawlerdegen/constants"
	"txcrawlerdegen/initializer"
)

var (
	HOST            string
	PORT            string
	ETHERSCAN_TOKEN string
	NODE_PROVIDER   string
)

func LoadVariables() {
	initializer.LoadEnv()
	HOST = constants.HOST
	PORT = os.Getenv("PORT")
	if len(PORT) == 0 {
		PORT = constants.PORT
	}
	// To make sure Pod does not spinup without a URL
	ETHERSCAN_TOKEN = os.Getenv("ETHERSCAN_TOKEN")
	if len(ETHERSCAN_TOKEN) == 0 {
		panic("No ETHERSCAN_TOKEN Provided. Panicking")
	}
	NODE_PROVIDER = os.Getenv("NODE_URL")
	if len(NODE_PROVIDER) == 0 {
		panic("No URL Provided. Panicking")
	}
}
