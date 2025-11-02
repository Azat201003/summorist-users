package main

import (
	"fmt"
	"github.com/Azat201003/summorist-users/internal/tokens"
)

func main() {
	var userId uint64
	fmt.Scan(&userId)
	r, err := tokens.GenerateToken(userId)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(r)
	}
}
