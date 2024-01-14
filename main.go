package main

import (
	"txcrawlerdegen/api"
	"txcrawlerdegen/internal/appconfig"
)

func main() {
	// First read all env variables
	appconfig.LoadVariables()
	// Then initialise API
	api.InitServer()
}
