package main

import (
	"github.com/Azat201003/summorist-users/internal/passwords"
	"fmt"
	"encoding/hex"
)

func main() {
	var password string
	fmt.Scan(&password)
	fmt.Println(hex.EncodeToString(passwords.Hash(password)))
}

