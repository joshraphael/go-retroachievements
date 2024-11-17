// Package getusersummary provides an example for a users summary info
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusersummary.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	games := 10
	achievements := 10
	resp, err := client.GetUserSummary(models.GetUserSummaryParameters{
		Username:          "jamiras",
		GamesCount:        &games,
		AchievementsCount: &achievements,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
