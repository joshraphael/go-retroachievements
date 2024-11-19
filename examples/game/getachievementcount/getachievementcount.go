// Package getachievementcount provides an example for getting the list of achievement IDs for a game.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getachievementcount.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetAchievementCount(models.GetAchievementCountParameters{
		GameID: 515,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
