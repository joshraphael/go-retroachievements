// Package getmostrecenttickets provides an example for getting ticket metadata information about the latest opened achievement tickets on RetroAchievements.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getmostrecenttickets.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	count := 100
	resp, err := client.GetMostRecentTickets(models.GetMostRecentTicketsParameters{
		Count: &count,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
