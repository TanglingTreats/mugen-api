package dotenv

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv() {
	fmt.Println("Loading .env")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
