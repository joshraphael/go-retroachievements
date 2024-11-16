// Package getuserrecentlyplayedgames provides an example for a users recently played games
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserrecentlyplayedgames.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	count := 10
	offset := 0
	resp, err := client.GetUserRecentlyPlayedGames(models.GetUserRecentlyPlayedGamesParameters{
		Username: "jamiras",
		Count:    &count,
		Offset:   &offset,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
