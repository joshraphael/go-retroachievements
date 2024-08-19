// Package getgameinfoanduserprogress provides an example for a users game info progress
package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements"
)

/*
Test script, add RA_API_KEY to your env and use `go run getgameinfoanduserprogress.go`
*/
func main() {
	secret := os.Getenv("RA_API_KEY")

	client := retroachievements.NewClient(secret)

	resp, err := client.GetGameInfoAndUserProgress("jamiras", 515, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}
