// Package getgameextended provides an example for getting a games info
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getgameextended.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGameExtended(293)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
