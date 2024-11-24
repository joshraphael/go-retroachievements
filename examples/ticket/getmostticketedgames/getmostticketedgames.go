// Package getmostticketedgames provides an example for getting the games on the site with the highest count of opened achievement tickets.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getmostticketedgames.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetMostTicketedGames(models.GetMostTicketedGamesParameters{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
