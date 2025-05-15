package env

import (
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
