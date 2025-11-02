package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Azat201003/summorist-users/internal/passwords"
)

func main() {
	var password string
	fmt.Scan(&password)
	fmt.Println(hex.EncodeToString(passwords.Hash(password)))
}
