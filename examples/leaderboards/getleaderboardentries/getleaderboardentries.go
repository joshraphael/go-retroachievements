// Package getleaderboardentries provides an example for getting a given leaderboards's entries.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getleaderboardentries.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetLeaderboardEntries(models.GetLeaderboardEntriesParameters{
		LeaderboardID: 104370,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
