// Package getusercompletedgames provides an example for getting a users completed games
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getusercompletedgames.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetUserCompletedGames(models.GetUserCompletedGamesParameters{
		Username: "jamiras",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}