// Package getgameleaderboards provides an example for getting a given games's list of leaderboards.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgameleaderboards.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGameLeaderboards(models.GetGameLeaderboardsParameters{
		GameID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
