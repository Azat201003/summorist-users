package main

import (
	"github.com/Azat201003/summorist-users/internal/server"
	//"github.com/Azat201003/summorist-users/internal/passwords"
	//"fmt"
)

func main() {
	server.StartServer("localhost", 8001)
}
