package main

import (
	"github.com/Azat201003/summorist-users/internal/server"
)

func main() {
	server.StartServer("localhost", 8001)
}
