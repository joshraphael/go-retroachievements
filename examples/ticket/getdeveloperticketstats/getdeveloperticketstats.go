// Package getdeveloperticketstats provides an example for getting ticket stats for a developer, targeted by that developer's site username.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getdeveloperticketstats.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetDeveloperTicketStats(models.GetDeveloperTicketStatsParameters{
		Username: "jamiras",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
