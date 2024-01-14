package initializer

import "github.com/joho/godotenv"

func LoadEnv() {
	_ = godotenv.Load("/Users/udayexp/Projects/degentask/.env")
}
