# txcrawlerdegen
Fetch USDC(default) or supported balance of user 

# Folder Structure

main.go -> entry point
api -> Folder contains everything related to API(middleware, handler, routes)
internal -> Contains env varibales initialiser, and internal logic. As more APIs come this becomes bigger
pkg -> External packages are added here. Ideally an abstract implementation should be exposed from here
constants -> All constants (no sensitive data) is present here

# CURL

curl --location --request GET 'http://127.0.0.1:5001/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "token_symbol": "USDC",
    "wallet_address": "0xDa9CE944a37d218c3302F6B82a094844C6ECEb17"
}'

# Current Flow

Above curl fetches the user balance
In case we do not send token_symbol, by default USDC is picked

If wrong symbol is sent, error displays all valid symbols to help users

Uses infura to communicate with B.C

# Steps to setup on local
 #TODO
