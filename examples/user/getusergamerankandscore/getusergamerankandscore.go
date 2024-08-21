// Package getusergamerankandscore get metadata about how a user has performed on a given game.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusergamerankandscore.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserGameRankAndScore("jamiras", 515)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
