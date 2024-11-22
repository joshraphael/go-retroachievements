// Package getachievementunlocks provides an example for getting a list of users who have earned an achievement
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementunlocks.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetAchievementUnlocks(models.GetAchievementUnlocksParameters{
		AchievementID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
