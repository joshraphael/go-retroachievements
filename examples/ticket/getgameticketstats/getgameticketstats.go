// Package getgameticketstats provides an example for getting ticket stats for a game, targeted by that game's unique ID.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgameticketstats.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	metadata := true
	resp, err := client.GetGameTicketStats(models.GetGameTicketStatsParameters{
		GameID:                1,
		IncludeTicketMetadata: &metadata,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
