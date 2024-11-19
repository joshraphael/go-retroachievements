// Package getgamehashes provides an example for getting a list of the games hashes
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
	"github.com/joshraphael/go-retroachievements/models"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgamehashes.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGameHashes(models.GetGameHashesParameters{
		GameID: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
