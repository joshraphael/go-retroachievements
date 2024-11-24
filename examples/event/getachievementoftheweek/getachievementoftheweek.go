// Package getachievementoftheweek provides an example for getting comprehensive metadata about the current Achievement of the Week.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementoftheweek.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetAchievementOfTheWeek(models.GetAchievementOfTheWeekParameters{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
