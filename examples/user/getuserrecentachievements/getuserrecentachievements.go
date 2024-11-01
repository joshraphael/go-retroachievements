// Package getuserrecentachievements provides an example for a users achievements in the last X minutes
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getuserrecentachievements.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	lookback := 1000
	resp, err := client.GetUserRecentAchievements(models.GetUserRecentAchievementsParameters{
		Username:        "Jamiras",
		LookbackMinutes: &lookback,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
