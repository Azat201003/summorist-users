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
		config.DBHost = os.Getenv("DBHOST")
		config.DBPort = os.Getenv("DBPORT")
		config.DBUser = os.Getenv("DBUSER")
		config.DBName = os.Getenv("DBNAME")
		config.DBPassword = os.Getenv("DBPASSWORD")
		isRead = true
	}
	return config
}
