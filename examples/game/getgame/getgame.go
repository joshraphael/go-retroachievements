// Package getgame provides an example for getting a games info
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for getting user profile. Add RA_API_KEY to your env and use `go run getgame.go`
*/
func main() {
	host := "https://retroachievements.org"
	secret := os.Getenv("RA_API_KEY")

	raClient := client.New(host, secret)

	resp, err := raClient.GetGame(293)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
