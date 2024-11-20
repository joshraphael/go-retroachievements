// Package getachievementdistribution provides an example for getting how many players have unlocked how many achievements for a game
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementdistribution.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	hardcore := true
	unofficial := false
	resp, err := client.GetAchievementDistribution(models.GetAchievementDistributionParameters{
		GameID:     1,
		Hardcore:   &hardcore,
		Unofficial: &unofficial,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
