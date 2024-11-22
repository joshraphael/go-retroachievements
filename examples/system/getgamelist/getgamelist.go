// getgamelist

// Package getgamelist provides an example for getting the complete list of games for a specified console on the site.
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgamelist.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGameList(models.GetGameListParameters{
		SystemID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
