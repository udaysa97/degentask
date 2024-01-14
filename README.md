# txcrawlerdegen
Fetch USDC(default) or supported balance of user 

# Folder Structure

main.go -> entry point
api -> Folder contains everything related to API(middleware, handler, routes)
internal -> Contains env varibales initialiser, and internal logic. As more APIs come this becomes bigger
pkg -> External packages are added here. Ideally an abstract implementation should be exposed from here
constants -> All constants (no sensitive data) is present here

# CURL

# Test functionailty created to fetch token blance based on symbol
curl --location --request GET 'http://127.0.0.1:5001/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token_symbol": "USDC",
    "wallet_address": "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
}'

# # functionailty created to fetch txs on eth (faced issues with providers)
curl --location --request GET 'http://127.0.0.1:5001/nft-balances' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wallet_address": "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
}'

# # functionailty created to fetch nft txs on wallet

curl --location --request GET 'http://127.0.0.1:5001/ethTxs' \
--header 'Content-Type: application/json' \
--data-raw '{
    "wallet_address": "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
}'

# Current Flow

Above curls fetches the user balances on token, eth txs and nft txs
In case we do not send token_symbol, by default USDC is picked (for test functionality API)

If wrong symbol is sent, error displays all valid symbols to help users

Uses infura and  to communicate with B.C

# Steps to setup on local
1) git clone https://github.com/udaysa97/degentask.git
2) cd degentask
3) go mod tidy
4) in .env file add your infura URL to connect to BSC chain and etherscan api key in the provided variables
(.env is not added to .gitignore to make it easier to understand where im accepting the secret values)
5) go run main.go (ideally should have been in cmd/main.go as per standards)