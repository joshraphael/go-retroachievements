package main

import (
	"fmt"
	"os"

	"github.com/joshraphael/go-retroachievements/client"
)

/*
Test script for various endpoints. Add RA_API_KEY to your env and use `go run sandbox.go`
*/
func main() {
	host := "https://retroachievements.org"
	secret := os.Getenv("RA_API_KEY")

	raClient := client.New(host, secret)

	resp, err := raClient.User.GetUserRecentAchievements("ChronoGear", 60)
	if err != nil {
		panic(err)
	}

	for i := range resp {
		fmt.Printf("%+v\n", resp[i])
	}
}
