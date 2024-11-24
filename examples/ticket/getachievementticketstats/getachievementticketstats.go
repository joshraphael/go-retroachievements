// Package getachievementticketstats provides an example for getting ticket stats for an achievement, targeted by that achievement's unique ID.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementticketstats.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetAchievementTicketStats(models.GetAchievementTicketStatsParameters{
		AchievementID: 284759,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
