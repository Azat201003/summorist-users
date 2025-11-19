package config

import (
	"os"
)

type Config struct {
	PrivateKey string
	PublicKey  string
	Host       string
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

var config Config
var isRead = false

func GetConfig() Config {
	if !isRead {
		// secrets.env
		config.PrivateKey = os.Getenv("PRIVATE_KEY")
		config.PublicKey = os.Getenv("PUBLIC_KEY")

		// config.env.*
		config.Host = os.Getenv("HOST")
		config.Port = os.Getenv("PORT")
		config.DBHost = os.Getenv("POSTGRES_HOST")
		config.DBPort = os.Getenv("POSTGRES_PORT")
		config.DBUser = os.Getenv("POSTGRES_USER")
		config.DBName = os.Getenv("POSTGRES_DB")
		config.DBPassword = os.Getenv("POSTGRES_PASSWORD")
		isRead = true
	}
	return config
}
