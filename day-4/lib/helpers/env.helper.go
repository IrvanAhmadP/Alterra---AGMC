package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

func init() {
	getConfig()
}

func getConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetDatabaseConfig() map[string]string {
	config := map[string]string{
		"DB_USERNAME": os.Getenv("DB_USERNAME"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("DB_NAME"),
	}
	return config
}

func GetAppEnvConfig() string {
	return os.Getenv("APP_ENV")
}

func GetSecretKeyConfig() string {
	return os.Getenv("SECRET_KEY")
}
