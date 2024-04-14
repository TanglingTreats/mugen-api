package dotenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnv(env string) {
	err := godotenv.Load(env)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
