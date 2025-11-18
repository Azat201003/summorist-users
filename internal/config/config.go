package config

import (
	"os"
)

type Config struct {
	PrivateKey string
	PublicKey string
	Host string
	Port string
}

var config Config
var isRead = false

func GetConfig() Config {
	if !isRead {
		config.PrivateKey = os.Getenv("PRIVATE_KEY")
		config.PublicKey = os.Getenv("PUBLIC_KEY")
		config.Host = os.Getenv("HOST")
		config.Port = os.Getenv("PORT")
		isRead = true
	}
	return config
}

