// Package getgame provides an example for getting a games info
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgame.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGame(293)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
