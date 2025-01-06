// Package getusergameleaderboards provides an example for getting a user's list of leaderboards for a given game.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusergameleaderboards.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserGameLeaderboards(models.GetUserGameLeaderboardsParameters{
		GameID:   583,
		Username: "joshraphael",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
