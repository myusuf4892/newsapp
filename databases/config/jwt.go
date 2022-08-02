package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func JWT() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env file")
	}
	SECRET_JWT := os.Getenv("SECRET_JWT")
	return SECRET_JWT
}
