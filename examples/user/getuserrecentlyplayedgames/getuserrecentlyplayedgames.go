// Package getuserrecentlyplayedgames provides an example for a users recently played games
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserrecentlyplayedgames.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserRecentlyPlayedGames("jamiras", 10, 0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
